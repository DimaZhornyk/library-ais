import {
  Component,
  createEffect,
  createMemo,
  createSignal,
  For,
  Show,
} from "solid-js";
import { styled } from "solid-styled-components";
import { DefaultAPI, ActionDTO } from "./api";
const Container = styled("table")`
  margin: auto;
`;
const Input = styled("input")`
  padding: 16px 32px;
  width: 200px;
`;

const Td = styled("td")`
  padding: 16px 32px;
`;

const CenteredRow = styled("div")`
  width: 100%;
  display: flex;
  justify-content: space-around;
`;

const Button = styled("button")`
  padding: 16px 32px;
  width: 70%;
  margin: auto;
`;
const tryParseNumber = (s: string): string | number => {
  const parsed = parseFloat(s);
  if (s.match("/^d+$/") == null) {
    return s;
  }
  return isNaN(parsed) ? s : parsed;
};
const setParams = (query: ActionDTO, params: Record<string, unknown>) => ({
  ...query,
  queries: query.queries.map((q) => ({
    ...q,
    params: Object.entries(q.params).reduce(
      (acc, [key]) => ({ ...acc, [key]: params[key] }),
      {}
    ),
  })),
});
export const QueryForm: Component<{
  query: ActionDTO;
  onResult: (result: Record<string, unknown>[]) => void;
}> = (props) => {
  const params = createMemo(() =>
    props.query.queries.reduce((acc, q) => ({ ...acc, ...q.params }), {})
  );
  const onSubmit = async (
    query: ActionDTO,
    params: Record<string, unknown>
  ) => {
    console.log({ params });
    await DefaultAPI.executeQuery(setParams(query, params))
      .then((r) => r ?? [])
      .then(props.onResult)
      .catch(() => alert("Bad Input"));
    setValue({});
  };
  const [value, setValue] = createSignal<Record<string, unknown>>({});
  createEffect(() => {
    if (Object.entries(params()).length == 0) {
      onSubmit(props.query, params());
    }
  });
  return (
    <Show when={Object.entries(params()).length > 0}>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          onSubmit(props.query, value());
        }}
      >
        <Container>
          <For each={Object.entries(params())}>
            {([name, type]) => (
              <tr>
                <Td>
                  <span>{name.split("_").join(" ")}</span>
                </Td>
                <td>
                  <Input
                    placeholder={name}
                    value={(value()[name] ?? "") + ""}
                    onInput={(e) =>
                      setValue({
                        ...value(),
                        [name]: tryParseNumber(e.currentTarget.value),
                      })
                    }
                    required
                  />
                </td>
              </tr>
            )}
          </For>
          <tr>
            <td colspan="2">
              <CenteredRow>
                <Button type="submit">Submit</Button>
              </CenteredRow>
            </td>
          </tr>
        </Container>
      </form>
    </Show>
  );
};

import {
  Accessor,
  Component,
  createEffect,
  createMemo,
  createResource,
  createSignal,
  For,
  Show,
  Suspense,
} from "solid-js";
import { styled } from "solid-styled-components";
import { DefaultAPI, EntityDTO, ActionDTO } from "./api";
import { QueryForm } from "./QueryForm";
const Table = styled("table")`
  color: white;
  width: 100%;
  border-collapse: collapse;
`;
const Td = styled("td")`
  text-align: center;
  padding: 16px 0;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
`;
const Th = styled("th")`
  text-align: center;
  padding: 16px 8px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
`;
const Tr = styled("tr")``;
const Title = styled("h3")`
  color: white;
  text-align: center;
`;
export const Entity: Component<{ entity: Accessor<EntityDTO> }> = (props) => {
  createEffect(() => {
    console.log({ props: props.entity().basicQuery });
  }, [props.entity]);
  const [executionRes, { refetch }] = createResource(
    props.entity,
    (e) =>
      DefaultAPI.executeQuery({
        queryName: "",
        queries: [{ text: e.basicQuery, params: {} }],
      }),
    { initialValue: [] }
  );

  const keys = createMemo(() => Object.keys(executionRes()[0] ?? {}));
  const values = createMemo(() =>
    executionRes().map((entry) => keys().map((key) => entry[key]))
  );
  createEffect(() => {
    console.log(executionRes());
  });
  return (
    <>
      {/*  <QueryForm query={props.query} onResult={setExecutionRes} /> */}
      <Show when={executionRes()}>
        <Table>
          <thead>
            <For each={keys()}>{(key) => <Th>{key}</Th>}</For>
          </thead>
          <tbody>
            <For each={values()}>
              {(value) => (
                <Tr>
                  <For each={value}>{(col) => <Td>{col + ""}</Td>}</For>
                </Tr>
              )}
            </For>
          </tbody>
        </Table>
      </Show>
      <Show when={executionRes().length === 0}>
        <Title>No Records</Title>
      </Show>
    </>
  );
};

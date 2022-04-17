import {
  Component,
  createEffect,
  createMemo,
  createResource,
  For,
  Suspense,
} from "solid-js";
import { styled } from "solid-styled-components";
import { DefaultAPI, QueryDTO } from "./api";
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
  padding: 16px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
`;
const Tr = styled("tr")``;
export const Query: Component<{ query: QueryDTO }> = (props) => {
  const [executionRes] = createResource(
    () => DefaultAPI.executeQuery(props.query),
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
    <Suspense>
      <Table>
        <thead>
          <For each={keys()}>{(key) => <Th>{key}</Th>}</For>
        </thead>
        <tbody>
          <For each={values()}>
            {(value) => (
              <Tr>
                <For each={value}>
                  {(col) => <Td>{JSON.stringify(col)}</Td>}
                </For>
              </Tr>
            )}
          </For>
        </tbody>
      </Table>
    </Suspense>
  );
};

import { Component, createMemo, createSignal, For, Show } from "solid-js";
import { styled } from "solid-styled-components";
import { ActionDTO } from "./api";
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
  padding: 16px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
`;
const Tr = styled("tr")``;
const Title = styled("h3")`
  color: white;
  text-align: center;
`;
const SubTitle = styled("h4")`
  color: white;
  text-align: center;
`;
export const Action: Component<{ action: ActionDTO }> = (props) => {
  console.log("Action");
  const [queryResult, setQueryResult] = createSignal<
    Record<string, unknown>[] | null
  >(null);
  const keys = createMemo(() => {
    const res = queryResult() ?? [];
    return Object.keys(res[0] ?? {});
  });
  const values = createMemo<unknown[][]>(() => {
    const res = queryResult() ?? [];
    return res.map((entry) => keys().map((key) => entry[key]));
  });
  return (
    <>
      <Title>{props.action.queryName}</Title>
      <Show when={queryResult() == null}>
        <QueryForm query={props.action} onResult={(r) => setQueryResult(r)} />
      </Show>
      <Show when={queryResult()}>
        <SubTitle>Completed!</SubTitle>
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
    </>
  );
};

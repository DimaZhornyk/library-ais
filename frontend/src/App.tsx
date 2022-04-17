import axios from "axios";
import {
  Component,
  createEffect,
  createMemo,
  createResource,
  createSignal,
  For,
  lazy,
  Show,
  Suspense,
} from "solid-js";
import { styled } from "solid-styled-components";
import { DefaultAPI, QueryDTO } from "./api";

import logo from "./logo.svg";
import { Query } from "./Query";
const MainContainer = styled("div")`
  height: 100vh;
  width: 100%;
  display: flex;
  flex-direction: row;
`;
const Card = styled("div")`
  backdrop-filter: blur(16px) saturate(180%);
  -webkit-backdrop-filter: blur(16px) saturate(180%);
  background-color: rgba(17, 25, 40, 0.75);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.125);
`;
const Sidebar = styled(Card)`
  width: 300px;
  margin: 20px;
  max-height: calc(100vh - 40px);
  overflow-x: scroll;
`;

const Main = styled(Card)`
  width: 800px;
  max-height: calc(100vh - 40px);
  height: calc(100vh - 40px);
  margin: 20px auto;
`;

const ListItem = styled("div")<{ clickable: boolean }>`
  cursor: ${(props) => (props.clickable ? "pointer" : "default")};
  border-bottom: 1px solid rgba(255, 255, 255, 0.1) !important;
  border-bottom-width: 1px !important;
  border-bottom-style: solid !important;
  border-bottom-color: rgba(255, 255, 255, 0.1) !important;
  padding: 16px 8px;
`;

const App: Component = () => {
  const [queries] = createResource(() => DefaultAPI.getQueries());
  const [selected, setSelected] = createSignal<QueryDTO>();
  const selectedParams = createMemo(() =>
    selected()?.queries.reduce((acc, q) => ({ ...acc, ...q.params }), {})
  );
  createEffect(() => {
    console.log(selectedParams());
  });
  return (
    <MainContainer>
      <Sidebar>
        <Suspense>
          <For each={queries()}>
            {(q) => (
              <ListItem clickable={true} onClick={() => setSelected(q)}>
                <p>{q.queryName}</p>
              </ListItem>
            )}
          </For>
        </Suspense>
      </Sidebar>
      <Main style={{ padding: "20px" }}>
        <Show when={selected()}>{(q) => <Query query={q} />}</Show>
      </Main>
    </MainContainer>
  );
};

export default App;

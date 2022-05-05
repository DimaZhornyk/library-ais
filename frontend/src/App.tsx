import axios from "axios";
import {
  Component,
  createResource,
  createSignal,
  For,
  Show,
  Suspense,
} from "solid-js";
import { styled } from "solid-styled-components";
import { DefaultAPI, EntityDTO, ActionDTO } from "./api";
import { Entity } from "./Entity";
import { Action } from "./Action";

import logo from "./logo.svg";
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
const NestedListItem = styled("div")<{ clickable: boolean }>`
  cursor: ${(props) => (props.clickable ? "pointer" : "default")};
  padding: 10px 8px;
  margin-left: 20px;
`;

const Link = styled("a")`
  color: white;
`;

const Nav = styled("div")`
  padding: 16px 8px;
  width: 100%;
  display: flex;
  justify-content: space-around;
`;

const App: Component = () => {
  const [queries] = createResource(() => DefaultAPI.getQueries());
  const [selected, setSelected] = createSignal<EntityDTO>();
  const [selectedAction, setSelectedAction] = createSignal<ActionDTO | null>();
  return (
    <>
      <Nav>
        <Link href="/admin">Admin</Link>
        <Link href="/librarian">Librarian</Link>
        <Link href="/reader">Reader</Link>
      </Nav>
      <MainContainer>
        <Sidebar>
          <Suspense>
            <For each={queries()}>
              {(q) => (
                <ListItem
                  clickable={true}
                  onClick={() => {
                    setSelected(q);
                    setSelectedAction(null);
                  }}
                >
                  <p>{q.entityName}</p>
                  <Show when={selected()?.basicQuery === q.basicQuery}>
                    <For each={q.actions}>
                      {(a) => (
                        <NestedListItem
                          clickable={true}
                          onClick={(e) => {
                            e.stopPropagation();
                            setSelectedAction(a);
                          }}
                        >
                          <p>{a.queryName}</p>
                        </NestedListItem>
                      )}
                    </For>
                  </Show>
                </ListItem>
              )}
            </For>
          </Suspense>
        </Sidebar>
        <Main style={{ padding: "20px" }}>
          <Show when={selected() && selectedAction() == null}>
            {() => <Entity entity={selected() as EntityDTO} />}
          </Show>
          <Show when={selectedAction()}>
            {(q) => {
              console.log({ q });
              return <Action action={q} />;
            }}
          </Show>
        </Main>
      </MainContainer>
    </>
  );
};

export default App;

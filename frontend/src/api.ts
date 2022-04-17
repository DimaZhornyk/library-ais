import axios from "axios";

export type QueryDTO = {
  queryName: string;
  queries: {
    text: string;
    params: Record<string, "string" | "int" | "float" | "date">;
  }[];
};
export class Api {
  constructor(
    private readonly _axios = axios.create({ baseURL: "http://localhost:9000" })
  ) {}
  public getQueries(): Promise<QueryDTO[]> {
    return this._axios.get<QueryDTO[]>("/queries").then((r) => r.data);
  }
  public executeQuery(query: QueryDTO): Promise<Record<string, unknown>[]> {
    return this._axios
      .post<Record<string, unknown>[]>("/query", query)
      .then((r) => r.data);
  }
}

export const DefaultAPI = new Api();

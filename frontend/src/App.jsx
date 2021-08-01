import React from "react";
import { Provider } from "urql";
import { client } from "./graphql";
import Header from "./Header";
import Users from "./Users";

function App() {
  return (
    <Provider value={client}>
      <Header />
      <Users />
    </Provider>
  );
}

export default App;

import React from "react";
import { Provider } from "urql";

import { client } from "./graphql";
import Explorer from "./components/explorer/Explorer";
import Header from "./components/Header";

function App() {
	return (
		<Provider value={client}>
			<Header />
			<Explorer />
		</Provider>
	);
}

export default App;

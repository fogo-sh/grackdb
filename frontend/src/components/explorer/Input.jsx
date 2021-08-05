import React, { useState } from "react";
import AceEditor from "react-ace";

import "ace-builds/src-noconflict/mode-graphqlschema";
import "ace-builds/src-noconflict/theme-twilight";

function Input({ query, setQuery }) {
	const [code, setCode] = useState(query);

	return (
		<div style={{ width: "40%" }}>
			<AceEditor
				defaultValue={query}
				mode="graphql"
				theme="twilight"
				onChange={(value) => setCode(value)}
				name="explorer-input"
				editorProps={{ $blockScrolling: true }}
				width="100%"
				height="30rem"
			/>
			<button
				className="w-full border border-white text-white"
				onClick={() => setQuery(code)}
			>
				Invoke Query
			</button>
		</div>
	);
}

export default Input;

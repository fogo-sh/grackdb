import React, { useState } from "react";
import { useQuery } from "urql";
import Input from "./Input";

import View from "./View";

const INITIAL_QUERY = `{
	projects {
		edges {
			node {
				id
				name
				startDate
				endDate
				technologies {
					id
					technology {
						name
						colour
					}
				}
				contributors {
					id
					user {
						id
						username
					}
				}
			}
		}
	}
}`;

function Explorer() {
	const [query, setQuery] = useState(INITIAL_QUERY);
	const [{ data }] = useQuery({ query });

	return (
		<div className="flex flex-row gap-4 pt-2">
			<Input query={query} setQuery={setQuery} />
			<View data={data} />
		</div>
	);
}

export default Explorer;

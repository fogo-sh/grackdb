import { useQuery } from "urql";

const query = /* GraphQL */ `
	{
		__schema {
			queryType {
				name
				kind
				fields {
					name
					type {
						name
					}
				}
			}
		}
	}
`;

const useQueryTypenameToField = () => {
	const [{ data }] = useQuery({ query });

	return (
		data &&
		Object.fromEntries(
			data.__schema.queryType.fields.map((field) => [
				field.type.name,
				field.name,
			])
		)
	);
};

export default useQueryTypenameToField;

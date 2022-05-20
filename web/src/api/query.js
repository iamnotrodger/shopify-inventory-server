export const buildQueryString = (queryObject) => {
	if (queryObject == null || Object.keys(queryObject).length === 0) {
		return '';
	}

	let queryString = '?';
	for (const [key, value] of Object.entries(queryObject)) {
		queryString += `${key}=${value}&`;
	}

	return queryString.slice(0, -1);
};

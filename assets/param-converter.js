function selectParamConverter(event, urlParamName) {
	let values = event.detail.value;
	let url = new URL(location.href);

	url.searchParams.delete(urlParamName);

	if (values.length === 0) {
		history.pushState(null, null, url.toString());

		return "";
	}

	url.searchParams.set(urlParamName, values.join(","));

	// history.pushState(null, null, `?${urlParamName}=${values}`);
	history.pushState(null, null, url.toString());

	return values.toString();
}

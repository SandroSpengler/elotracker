function selectParamConverter(event, urlParamName) {
    let values = event.detail.value;
    let url = new URL(location.href);

    url.searchParams.delete(urlParamName);

    if (Array.isArray(values)) {
        if (values.length > 0) {
            url.searchParams.set(urlParamName, values.join(","));
        }
    }

    if (event.detail && typeof event.detail === 'object' && !Array.isArray(values)) {
        url.searchParams.set(urlParamName, values)
    }

    let params = Object.fromEntries(url.searchParams.entries());

    history.pushState(null, null, url.toString());

    return params;
}


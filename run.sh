runType=$1;

if [[ $runType == "go" ]]; then
     $(go env GOPATH)/bin/air
fi

if [[ $runType == "tailwind" ]]; then
    npx tailwindcss -i styles/tailwind.css -o assets/main.css --watch
fi


# live reload templ files
# templ generate --watch --proxy="http://localhost:5555" --proxyport="5555" --cmd="go run ."
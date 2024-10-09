runType=$1;

if [[ $runType == "go" ]]; then
     $(go env GOPATH)/bin/air
fi

if [[ $runType == "tailwind" ]]; then
    npx tailwindcss -i src/styles/tailwind.css -o assets/main.css --watch
fi

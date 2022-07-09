<h1 align="center">IT'S NOT WORKING ANYMORE</h1>

------

<h2 align="center"> >> LlamaLauncher << </h2> 

<h4 align="center">Simple Fortnite Launcher with DLL-Injection ability</a></h4>
<h6 align="center">inspired by <a href="https://github.com/Londiuh/FortniteLauncher">Londiuh's Launcher</a></h6>

------

## How to use it?

  - Download the latest release ( or [build it](#how-to-build) yourself )
  - Open Fiddler (if you want to debug the HTTP requests)
  - And just start `llamalauncher.exe`!

<details>
  <summary>Add new account:</summary>
  
  ![HowToAddAccount](https://user-images.githubusercontent.com/60406325/153208329-b2147d04-741c-4599-9883-2eea2558386a.gif)
</details>

<details>
  <summary>Use existing accounts:</summary>
  
  ![HowToUseAccount](https://user-images.githubusercontent.com/60406325/153207671-e8e2bf62-28c5-4ff0-9cbb-aaca034c9600.gif)
</details>

------

## How to build?

### Prerequisites:
  - Having [Golang](https://go.dev/dl/) installed
  - Having [Entgo](https://entgo.io/) installed
  - Having the `gcc` compiler in your PATH

### Build commands:
```bash
$ git clone https://github.com/MR-AliHaashemi/LlamaLauncher
$ cd LlamaLauncher
$ ent generate ./pkg/database/schema/
$ go mod tidy
$ go build .
```

------

## Special thanks to:
  - [er-azh](https://github.com/er-azh) for [egmanifest](https://github.com/er-azh/egmanifest) and the dll injection code
  - [FortniteModding](https://github.com/FortniteModding) Organization for [PlataniumV2](https://github.com/FortniteModding/PlataniumV2)

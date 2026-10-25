# 2048 Froggy

An implementation of the 2048 game in Go using the [gofroggy](https://github.com/Maheshivara/gofroggy) framework, targeting the SF3000 handheld and other devices from the same family.

## Building

### Docker

1. Copy `.env.example` to `.env` and adjust the values to match your setup.
2. Build using Docker Compose:
   ```sh
   docker compose up --build
   ```
3. The compiled MIPS binary will be available at `./out/2048`.

4. To package the game as a port for [TreeFrogUI](https://github.com/tzubertowski/TreeFrogUI), run:

   ```sh
   chmod +x ./tools/package.sh && ./tools/package.sh
   ```

   The package will be generated at `./out/2048.zip`.

### Toolchain

1. Download and install a compatible MIPS toolchain, such as the [SF3000 Toolchain](https://github.com/game-de-it/sf3000/releases/tag/sf3000_toolchain_v0.1).
2. Configure `docker/mips-gcc` to match your toolchain installation, mainly the compiler and sysroot paths.
3. Install Go and configure the environment for MIPS cross-compilation:
   ```sh
   export CGO_ENABLED=1
   export GOOS=linux
   export GOARCH=mipsle
   export GOMIPS=hardfloat
   export CC=/path/to/your/mips-gcc
   ```
4. Build the game:

   ```sh
   go build -o ./out/2048 .
   ```

5. Optionally, package the game for TreeFrogUI using the same packaging command described above.

## Credits

- [gofroggy](https://github.com/Maheshivara/gofroggy) - Go framework for SF3000 devices.
- [TreeFrogUI](https://github.com/tzubertowski/TreeFrogUI) - Custom frontend for supported handhelds.
- [game-de-it](https://github.com/game-de-it/sf3000) - MIPS cross-compilation toolchain.

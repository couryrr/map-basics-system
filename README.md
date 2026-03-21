**Factory RTS style top-down game in Go + raylib-go**

---

## Build

```bash
make windows   # cross-compile and deploy to Windows
make linux     # local Linux build
```

Deploys to `/mnt/c/Users/Coury/Devlopment/map-basics.exe` (WSL path).

---

## Architecture

**Rendering:** World → Camera2D → virtual render texture (960×540) → letterboxed screen blit

**Events:** Callback pub/sub broker (`system/pubsub`). Topics are typed string constants owned by the publishing package. Input controller converts screen→virtual coords and publishes; subsystems subscribe.

**UI framework** (`system/ui/`): Container-based element tree. `Container` holds children, a `Style` (padding, gap, border, offset, etc.), and a layout engine. Layout is applied at `AddChild` time — static, not per-frame. Supported layouts: `LayoutHorizontal`, `LayoutVertical`, `LayoutGrid`. Style is composed with `With*` functional options (`WithLayout`, `WithPadding`, `WithBorder`, etc.).

**World:** Infinite terrain via 3-octave FBM Perlin noise (cached PNGs). 32×32 tile chunks, 5-chunk render radius around camera.

**Item registry:** `GameItem` structs loaded from `assets/directory.json` at startup. Accessed via `RegistryState` interface.

---

## Package map

| Package | Role |
|---|---|
| `config/` | Virtual resolution constants |
| `system/pubsub/` | Broker, Topic, Message |
| `system/controller/` | Input → virtual coords → broker |
| `system/renderer/` | RenderContext, viewport/letterbox |
| `system/ui/` | UI framework, overlay, hotbar, registry |
| `entity/player/` | Player state, Hotbar |
| `world/` | Terrain, chunk rendering, item registry |
| `game.go` | Composition root, `DrawContext` impl |
| `main.go` | Window init, broker wiring, game loop |

---

## Open

- Buildings and entity placement not yet implemented
- World save format undecided
- UI mouse interaction (hit-testing) stubbed, not wired
- Draw order / Z-sorting deferred

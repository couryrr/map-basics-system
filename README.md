**Factory RTS style top-down game in Go + raylib-go**

---

## Rendering pipeline

World space -> Camera transform -> Virtual render texture -> Screen blit

- Virtual resolution: 960x540
- Tile size: 32x32 pixels
- `VirtualWidth`/`VirtualHeight` are constants in the `config` package
- Blit uses `DrawTexturePro` with a letterbox rect to handle arbitrary window sizes

---

## Coordinate transforms

- Mouse: screen -> virtual (manual scale/offset) -> world (`rl.GetScreenToWorld2D`)
- UI collision checks stop at virtual space
- World collision checks go all the way to world space
- `system/controller` does the screen -> virtual conversion once and publishes virtual coords

---

## World and terrain

- FIXME: Terrain uses fBm (3 octave layered Perlin noise, cached as PNG) better lib needed
- Active chunk area: 5 chunks in each direction around camera
- The world is effectively infinite, Perlin sampling wraps via modulo

---

## Project structure

| Package | Contents |
|---|---|
| `config/` | VirtualWidth, VirtualHeight |
| `system/camera/` | Camera modes, rotation-aware movement |
| `system/controller/` | Raw input, screen->virtual coord conversion, topic publishing |
| `system/pubsub/` | Broker, Message, Topic types |
| `system/renderer/` | RenderContext, viewport/letterbox management |
| `system/setting/` | Screen size settings |
| `system/ui/` | InGameOverlay, HotbarElement, HotbarState, DrawContext interfaces |
| `entity/player/` | Player, Hotbar |
| `world/` | Terrain generation, chunk rendering |
| `game.go` | Game struct, composition root |
| `main.go` | Window init, broker wiring, game loop |

---

## Event bus

- Callback-based pub/sub in `system/pubsub`, single-threaded with the game loop (single-threaded atm)
- `Topic` is a typed `string` alias, constants defined in the package that owns the event
- `Broker.Register(topic, callback)` and `Broker.Send(topic, message)` are the full API
- Message payload is `any`, receivers type-assert on receipt (fixme not sure about this any)

---

## Player and hotbar

- Player owns Hotbar, single source of truth for slot contents and active slot
- `Hotbar` satisfies the `ui.HotbarState` interface (`SlotItem(i)`, `GetActiveSlot()`) used by the draw path
- `HotbarElement` (UI) owns only the slot bounds rectangles
- Hotbar interaction uses a typed `HotbarAction` string with defined constants (`hover`, `leave`)
- Interaction flow: cursor moved -> `InGameOverlay.CheckIntersection` -> `TopicUiHotbarInteraction` -> `Player.HandleHotbarInteraction` dispatches on action

---

## Input / controller

- `system/controller` reads raw input, converts screen -> virtual coords, publishes to broker
- `TopicInputCursorMoved` only fires when `rl.GetMouseDelta()` is non-zero
- UI elements return an `InteractionResult{Topic, Message}`, `InGameOverlay` publishes if non-nil

---

## Open decisions

- Art approach: placeholder rects, free asset pack
- Buildings and entities: not yet implemented
- Directory/Registry (item catalog from JSON): not yet implemented
- World save format: not yet decided
- Layer/Z sorting for draw order: not yet implemented
- Spatial partitioning for world-space click interaction: deferred until needed

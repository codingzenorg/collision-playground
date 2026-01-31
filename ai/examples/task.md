Task:
<one concrete goal>

Constraints:
<optional, only if special>

Done when:
<observable stop condition>
---
Task:
Render 20 circles with random positions and velocities.

Done when:
Circles move smoothly and bounce on window edges.


---
# Task

## Summary
<one sentence>

## Done when
- [ ] <observable condition 1>
- [ ] <observable condition 2>

## Notes (optional)
- <any constraints or non-goals>
---
# Task

## Summary
Add basic circle movement using velocity and frame delta.

## Done when
- [ ] Circles move smoothly on screen
- [ ] Movement uses velocity vectors (no hardcoded steps)
- [ ] Frame rate remains stable
- [ ] Works in native and WASM builds

## Non-goals
- No circle–circle collision
- No physics abstraction layer
- No refactors outside rendering and update loop

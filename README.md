# Tempest

<img alt="Github license badge" src="https://img.shields.io/github/license/Amatsagu/tempest" />
<img alt="Maintenance badge" src="https://img.shields.io/maintenance/yes/2024" />

> A simple, robust framework for creating discord applications in typescript (for deno runtime).

- Easily scalable!
- No cache by default,
- Fully typed,

## Supported parts

- [x] Webhook web server for receiving incoming payloads
- [x] (Slash) Command handler (both normal & sub commands)
- [x] Buttons
- [x] Creating/Editing/Deleting/Crossposting regular messages
- [x] REST handler with built-in rate limit protection
- [x] Followes camelCase (all Discord's snakecase payloads follows JS/TS standards)
- [x] Data compression to lower memory footprint (ids are turn into bigints & some codes into hashes)
- [x] Helpful error messages when creating interactions
- [x] Select menus
- [ ] User/Text messages commands
- [ ] Modals

❌ Multi-language support<br />
❌ Sending/receiving files

Elements with red cross won't be supported in this library because they would make code far more complex and barely
anyone would use it. You can listen to raw payloads and freely use REST API if you want so if you feel you need those -
add them for yourself.

## Special parts

- [x] Advanced cooldown system
- [x] Button menus handler via `client#listenButtons`

## Performance

Tempest is interaction focused library for Discord apps. Thanks to not using gateway, we gain a lot of efficiency. How
much?

Deno uses Rust's Hyper crate for dealing with networking
([benchmark](https://deno.land/benchmarks#http-server-throughput)). Average deno http server can handle around
`40K req/sec on Windows` and about `70K req/sec on Linux`. Assuming you use linux server - your app would need
(approximately) `~300K discord guilds` to hit throughput issues. That's efficiency of `~120 gateway shards`! On top of
that - single webhook will likely take far less resources than process with 60 ws sockets. Additionally - scalling
discord apps is super easy. Just spawn new mirror process and link it with for example nginx's balanceloader. Scalling
gateway based bot can be a nightmare.

All of that cost you just a bit higher average ping and of course Discord apps are still a bit limited in functionality.
Pick your poison :)

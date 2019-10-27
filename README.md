<!-- portsmith — protocol cartographer's field manual -->
<p align="center">
  <img src="docs/assets/hero.svg" alt="portsmith — a protocol cartographer's field manual: an oscilloscope tracing packets through capture, normalize, infer and replay" width="100%">
</p>

<h1 align="center">portsmith</h1>

<p align="center">
  <strong>Protocol discovery &amp; replay for line-based network traffic.</strong><br>
  <em>Put an unknown wire protocol on the bench, watch its waveform, chart its schema, and replay it.</em>
</p>

<p align="center">
  <code>capture</code> &#8594; <code>normalize</code> &#8594; <code>infer</code> &#8594; <code>replay</code> &#183;
  <code>Go</code> + <code>Rust</code> &#183; standard libraries only &#183; MIT
</p>

---

## Field manual

You have a TCP service with no spec — an undocumented daemon, a vendor box that
speaks its own dialect, a Redis-flavoured thing you want to regression-test. You
have a socket, not a document.


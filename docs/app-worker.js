const cacheName = "app-" + "d4334907d2d35847f5e635382bdb1f3d1e78d082";

self.addEventListener("install", event => {
  console.log("installing app worker d4334907d2d35847f5e635382bdb1f3d1e78d082");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "",
        "/",
        "/app.css",
        "/app.js",
        "/manifest.webmanifest",
        "/wasm_exec.js",
        "/web/app.wasm",
        "/web/lofimusic.css",
        "/web/logo.png",
        "https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
        
      ]);
    })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker d4334907d2d35847f5e635382bdb1f3d1e78d082 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});

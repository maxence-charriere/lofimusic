const cacheName = "app-" + "c31325d1a3f83d4431479cdc03e3e9bb539d26dc";

self.addEventListener("install", event => {
  console.log("installing app worker c31325d1a3f83d4431479cdc03e3e9bb539d26dc");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "/",
        "/app.css",
        "/app.js",
        "/manifest.json",
        "/wasm_exec.js",
        "/web/app.wasm",
        "/web/lofimusic.css",
        "/web/logo.png",
        "https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
        "https://fonts.googleapis.com/css2?family=Roboto&display=swap",
        
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
  console.log("app worker c31325d1a3f83d4431479cdc03e3e9bb539d26dc is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});

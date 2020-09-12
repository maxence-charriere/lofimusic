const cacheName = "app-" + "7c8c19537304800e79af75859aa37d4ec9930d32";

self.addEventListener("install", event => {
  console.log("installing app worker 7c8c19537304800e79af75859aa37d4ec9930d32");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "/lofimusic",
        "/lofimusic/app.css",
        "/lofimusic/app.js",
        "/lofimusic/manifest.json",
        "/lofimusic/wasm_exec.js",
        "/lofimusic/web/app.wasm",
        "/lofimusic/web/lofimusic.css",
        "https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
        "https://fonts.googleapis.com/css2?family=Roboto&display=swap",
        "https://storage.googleapis.com/murlok-github/icon-192.png",
        "https://storage.googleapis.com/murlok-github/icon-512.png",
        
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
  console.log("app worker 7c8c19537304800e79af75859aa37d4ec9930d32 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});

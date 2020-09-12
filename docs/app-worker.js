const cacheName = "app-" + "4d9d30a3b0d73a2193df5a2bdce8951376d4bc92";

self.addEventListener("install", event => {
  console.log("installing app worker 4d9d30a3b0d73a2193df5a2bdce8951376d4bc92");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "/",
        "//app.css",
        "//app.js",
        "//manifest.json",
        "//wasm_exec.js",
        "//web/app.wasm",
        "//web/lofimusic.css",
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
  console.log("app worker 4d9d30a3b0d73a2193df5a2bdce8951376d4bc92 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});

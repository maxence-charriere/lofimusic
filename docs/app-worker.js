const cacheName = "app-" + "d236fd2df359eac46fb11634c3d07eb1e00b1eda";

self.addEventListener("install", event => {
  console.log("installing app worker d236fd2df359eac46fb11634c3d07eb1e00b1eda");
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
  console.log("app worker d236fd2df359eac46fb11634c3d07eb1e00b1eda is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});

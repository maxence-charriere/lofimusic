const cacheName = "app-" + "266ef15c60a7f8062301ac5a55546b825540d317";

self.addEventListener("install", event => {
  console.log("installing app worker 266ef15c60a7f8062301ac5a55546b825540d317");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
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
      }).
      then(() => {
        self.skipWaiting();
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
  console.log("app worker 266ef15c60a7f8062301ac5a55546b825540d317 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});

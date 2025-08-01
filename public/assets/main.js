(function () {
  animateElementsInView();
  document.addEventListener("scroll", throttle(animateElementsInView, 100));
  document.addEventListener("scroll", hideDownIcon);

  function hideDownIcon() {
    if (window.scrollY >= window.innerHeight) {
      const el = document.querySelector(".down-icon");
      el?.remove();
      document.removeEventListener("scroll", hideDownIcon);
    }
  }

  function throttle(fn, delay) {
    let lastCall = 0;
    return function (...args) {
      const now = new Date().getTime();
      if (now - lastCall < delay) return;
      lastCall = now;
      return fn(...args);
    };
  }

  function animateElementsInView() {
    const animatedElements = document.querySelectorAll(".animate-in");
    animatedElements.forEach((el) => {
      const isInView =
        el.parentElement?.getBoundingClientRect().top < window.innerHeight;
      const isRunning = el.style["animation-play-state"] === "running";
      if (!isInView || isRunning) return;

      el.style["animation-delay"] = `${offsetMs(el)}ms`;
      el.style["animation-play-state"] = "running";
    });
  }

  function offsetMs(el, maxDelay = 500) {
    const offsetFromTop = el.getBoundingClientRect().top;
    const windowHeight = window.innerHeight;
    return maxDelay * (offsetFromTop / windowHeight);
  }
})();

;(function () {
  animateElementsInView()
  document.on('scroll', throttle(animateElementsInView, 20))
  document.on('scroll', hideDownIcon)

  function hideDownIcon() {
    if (window.scrollY >= window.innerHeight) {
      $('.down-icon')?.remove()
      document.removeEventListener('scroll', hideDownIcon)
    }
  }

  function throttle(fn, delay) {
    let lastCall = 0
    return function (...args) {
      const now = new Date().getTime()
      if (now - lastCall < delay) return
      lastCall = now
      return fn(...args)
    }
  }

  function animateElementsInView() {
    $$('.animate-in').forEach(el => {
      const isInView = el.parentElement?.getBoundingClientRect().top < window.innerHeight
      const isRunning = el.style['animation-play-state'] === 'running'
      if (!isInView || isRunning) return

      el.style['animation-delay'] = `${offsetMs(el)}ms`
      el.style['animation-play-state'] = 'running'
    })
  }

  function offsetMs(el, maxDelay = 500) {
    const offsetFromTop = el.getBoundingClientRect().top
    const windowHeight = window.innerHeight
    return maxDelay * (offsetFromTop / windowHeight)
  }
})()

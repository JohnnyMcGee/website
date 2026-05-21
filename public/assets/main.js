;(function () {
  const contact = $('#contact')
  contact.on('submit', handleContactSubmit)
  contact.on('input', handleContactInput)

  animateElementsInView()
  document.on('scroll', throttle(animateElementsInView, 20))
  document.on('scroll', hideDownIcon)

  const fadeElements = $$('.fade-out-on-scroll')
  fadeOutOnScroll()
  document.on('scroll', fadeOutOnScroll)

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

  function handleContactInput(event) {
    const form = event.target.form
    const formData = new FormData(form)
    $('#contact [type="submit"]').toggleAttribute(
      'disabled',
      !formData.get('name') || !formData.get('email') || !formData.get('message')
    )
  }
  Input
  function handleContactSubmit(event) {
    event.preventDefault()
    const form = event.target
    const data = Object.fromEntries(new FormData(form).entries())

    fetch(form.action, {
      method: form.method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    })
      .then(response => {
        if (response.ok) {
          alert('Message sent successfully!')
          form.reset()
        } else {
          alert('Failed to send message. Please try again later.')
        }
      })
      .catch(() => {
        alert('An error occurred. Please try again later.')
      })
  }

  function fadeOutOnScroll() {
    const scrollY = window.scrollY
    const windowHeight = window.innerHeight

    fadeElements.forEach(el => {
      el.style.opacity = Math.max(0, 1 - scrollY / windowHeight)
    })
  }
})()

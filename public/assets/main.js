;(function () {
  animateElementsInView()
  document.on('scroll', throttle(animateElementsInView, 20))
  document.on('scroll', hideDownIcon)
  window.initSwiper = initSwiper

  document.on('DOMContentLoaded', () => {
    const contact = $('#contact')
    contact.on('submit', handleContactSubmit)
    contact.on('input', handleContactInput)

    const fadeElements = $$('.fade-out-on-scroll')
    fadeOutOnScroll(fadeElements)
    document.on('scroll', () => fadeOutOnScroll(fadeElements))
  })

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

  async function handleContactSubmit(event) {
    event.preventDefault()
    const form = event.target
    const data = Object.fromEntries(new FormData(form).entries())
    const { ok, error } = await sendMessage(data)

    if (!ok) {
      console.error('Contact form submission failed:', error)
      $('#contact button[type="submit"]')?.insertAdjacentElement(
        'beforebegin',
        createErrorEl(error)
      )
      return
    }

    for (const el of $$('#contact .error')) {
      el.remove()
    }
    form.classList.add('success')
    form.reset()
  }

  function createErrorEl(message) {
    const el = document.createElement('div')
    el.className = 'error'
    el.textContent = message
    return el
  }

  async function sendMessage(data) {
    try {
      const url = '/api/contact'
      const res = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
      })
      const body = await res.json()
      return body
    } catch (e) {
      console.error('Error sending message:', e)
      return { ok: false, error: 'An error occurred. Please try again later.' }
    }
  }

  function fadeOutOnScroll(els) {
    const scrollY = window.scrollY
    const windowHeight = window.innerHeight

    els.forEach(el => {
      el.style.opacity = Math.max(0, 1 - scrollY / windowHeight)
    })
  }

  function initSwiper() {
    console.log('Initializing swiper...')
    const swiper = new window.Swiper('.swiper', {
      // Optional parameters
      loop: true,

      // If we need pagination
      pagination: {
        el: '.swiper-pagination'
      },

      // Navigation arrows
      navigation: {
        nextEl: '.swiper-button-next',
        prevEl: '.swiper-button-prev'
      },

      // And if we need scrollbar
      scrollbar: {
        el: '.swiper-scrollbar'
      }
    })
  }
})()

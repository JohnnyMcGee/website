(function () {
    const menuButton = document.getElementById("menu-button");
    menuButton?.addEventListener("click", handleClickMenu);

    animateElementsInView();
    document.addEventListener("scroll", throttle(animateElementsInView, 100));
    document.addEventListener("scroll", headerScrollListener());

    function throttle(fn, delay) {
        let lastCall = 0;
        return function (...args) {
            const now = new Date().getTime();
            if (now - lastCall < delay) return;
            lastCall = now;
            return fn(...args);
        };
    }

    function headerScrollListener() {
        let referenceScroll = window.scrollY;
        let lastScroll = window.scrollY;

        return () => {
            const header = document.querySelector("header");

            const dy = window.scrollY - referenceScroll;

            if (dy > 100) {
                header.classList.add("header--scrolled");
            }

            if (dy < -20 || window.scrollY < 100) {
                header.classList.remove("header--scrolled");
            }

            if ((window.scrollY < lastScroll && window.scrollY > referenceScroll) || (window.scrollY > lastScroll && window.scrollY < referenceScroll)) {
                referenceScroll = window.scrollY;
            }

            lastScroll = window.scrollY;
        };
    }

    function animateElementsInView() {
        const animatedElements = document.querySelectorAll(".animate-in");
        animatedElements.forEach((el) => {
            const isInView = el.parentElement?.getBoundingClientRect().top < window.innerHeight;
            const isRunning = el.style['animation-play-state'] === 'running';
            if (!isInView || isRunning) return;

            el.style['animation-delay'] = `${offsetMs(el)}ms`;
            el.style['animation-play-state'] = 'running';
        });
    }

    function offsetMs(el, maxDelay = 500) {
        const offsetFromTop = el.getBoundingClientRect().top;
        const windowHeight = window.innerHeight;
        return maxDelay * (offsetFromTop / windowHeight);
    }

    function handleClickMenu(event) {
        const button = event.currentTarget;

        if (!button) {
            throw new Error("Menu button not found");
        }

        const drawer = document.getElementById("menu-drawer");

        if (!drawer) {
            throw new Error("Menu drawer not found");
        }

        const isOpen = drawer.classList.contains("drawer--open");

        function closeOnScroll() {
            closeMenu(button, drawer);
            document.removeEventListener("scroll", closeOnScroll);
        }

        if (isOpen) {
            closeMenu(button, drawer);
            document.removeEventListener("scroll", closeOnScroll);

            return;
        }

        openMenu(button, drawer);
        document.addEventListener("scroll", closeOnScroll);
    }

    function closeMenu(button, drawer) {
        button.classList.remove("hamburger--open");
        drawer.classList.remove("drawer--open");
        drawer.setAttribute("aria-hidden", "true");
    }

    function openMenu(button, drawer) {
        button.classList.add("hamburger--open");
        drawer.classList.add("drawer--open");
        drawer.removeAttribute("aria-hidden");
    }
}
)();

[...document.querySelectorAll('*')]
    .flatMap(el => [...el.classList])
    .filter((cls, _, arr) => {
        return ![...document.styleSheets]
            .flatMap(s => { try { return [...s.cssRules] } catch { return [] } })
            .some(r => r.selectorText?.includes('.' + cls))
    })
    .filter((v, i, a) => a.indexOf(v) === i)
    .sort()
    .forEach(cls => console.warn('Not in CSS:', cls))
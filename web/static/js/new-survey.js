function toggleSeaGroup(group) {
    const checkbox = document.getElementById(group + '-enabled');
    const selects = document.getElementById(group + '-group');
    const select = document.getElementById(group === 'wave' ? 'sea-condition' : 'ice-condition');

    if (checkbox.checked) {
        selects.classList.remove('sea-selects--disabled');
        select.disabled = false;
        if (group === 'wave') updateSeaBadge();
        if (group === 'ice') updateIceBadge();
    } else {
        selects.classList.add('sea-selects--disabled');
        select.disabled = true;
    }
}

function updateSeaBadge() {
    const map = {
        "< 0.1m": "Calm",
        "0.1-0.5m": "Smooth",
        "0.5-1.25m": "Slight",
        "1.25-2.5m": "Moderate",
        "2.5-4.0m": "Rough"
    };
    const sel = document.getElementById('sea-condition');
    const badge = document.getElementById('sea-badge');
    const val = sel.value;
    badge.textContent = '';
    const dot = document.createElement('span');
    dot.className = 'sea-badge-dot';
    badge.appendChild(dot);
    badge.appendChild(document.createTextNode(map[val] || '—'));
    badge.className = 'sea-badge ' + (val || '');
}

function updateIceBadge() {
    const sel = document.getElementById('ice-condition');
    const badge = document.getElementById('ice-badge');
    const val = sel.value;
    badge.textContent = '';
    const dot = document.createElement('span');
    dot.className = 'sea-badge-dot';
    badge.appendChild(dot);
    badge.appendChild(document.createTextNode(val || '—'));
    badge.className = 'sea-badge';
}

// ── Sidebar active on scroll ──────────────────────────────────
const sections = ['section-survey', 'section-vessel', 'section-marks', 'section-calc'];
const navMap = {
    'section-survey': 'nav-survey',
    'section-vessel': 'nav-vessel',
    'section-marks': 'nav-marks',
    'section-calc': 'nav-calc'
};

window.addEventListener('scroll', () => {
    let current = sections[0];
    sections.forEach(id => {
        const el = document.getElementById(id);
        if (el && el.getBoundingClientRect().top < 80) current = id;
    });
    document.querySelectorAll('.sidebar-nav-item').forEach(i => i.classList.remove('is-active'));
    document.getElementById(navMap[current]).classList.add('is-active');
});

// ── LBP / 2 calc hint ─────────────────────────────────────────
function updateCalcHints() {
    const lbp = parseFloat(document.getElementById('lbp').value);
    document.getElementById('lbp-half').textContent = isNaN(lbp) ? '—' : (lbp / 2).toFixed(2) + ' m';
}

// ── FWA & 0.5% DWT ───────────────────────────────────────────
function updateFWA() {
    const dwt = parseFloat(document.getElementById('summer-dwt').value);
    const tpc = parseFloat(document.getElementById('summer-tpc').value);
    document.getElementById('half-pct-dwt').textContent = isNaN(dwt) ? '—' : (dwt * 0.005).toFixed(3) + ' MT';
    if (!isNaN(dwt) && !isNaN(tpc) && tpc > 0) {
        document.getElementById('fwa').textContent = (dwt / (4 * tpc)).toFixed(1) + ' mm';
    } else {
        document.getElementById('fwa').textContent = '—';
    }
}

// ── Depth check ───────────────────────────────────────────────
function checkDepth() {
    const depth = parseFloat(document.getElementById('depth').value);
    const draft = parseFloat(document.getElementById('summer-draft').value);
    const fb = parseFloat(document.getElementById('summer-freeboard').value);
    const warn = document.getElementById('depth-warning');
    const check = document.getElementById('depth-check');

    if (!isNaN(draft) && !isNaN(fb)) {
        const expected = (draft + fb).toFixed(3);
        check.textContent = expected + ' m';
        if (!isNaN(depth) && Math.abs(depth - parseFloat(expected)) > 0.001) {
            warn.style.display = 'flex';
            check.style.color = 'var(--danger)';
        } else {
            warn.style.display = 'none';
            check.style.color = 'var(--text-calc)';
        }
    } else {
        check.textContent = '—';
    }
}

function toggleSeaGroup(group) {
    const other = group === 'wave' ? 'ice' : 'wave';

    // switch on selected
    document.getElementById(group + '-group').classList.remove('sea-selects--disabled');
    document.getElementById(group === 'wave' ? 'sea-condition' : 'ice-condition').disabled = false;

    // switch on other
    document.getElementById(other + '-group').classList.add('sea-selects--disabled');
    document.getElementById(other === 'wave' ? 'sea-condition' : 'ice-condition').disabled = true;

    // update badge
    if (group === 'wave') updateSeaBadge();
    else updateIceBadge();
}

function updateSeaBadge() {
    const labelMap = {
        "< 0.1m": "Calm",
        "0.1-0.5m": "Smooth",
        "0.5-1.25m": "Slight",
        "1.25-2.5m": "Moderate",
        "2.5-4.0m": "Rough"
    };
    const classMap = {
        "< 0.1m": "calm",
        "0.1-0.5m": "smooth",
        "0.5-1.25m": "slight",
        "1.25-2.5m": "moderate",
        "2.5-4.0m": "rough"
    };
    const sel = document.getElementById('sea-condition');
    const badge = document.getElementById('sea-badge');
    const val = sel.value;
    badge.textContent = '';
    const dot = document.createElement('span');
    dot.className = 'sea-badge-dot';
    badge.appendChild(dot);
    badge.appendChild(document.createTextNode(labelMap[val] || '—'));
    badge.className = 'sea-badge ' + (classMap[val] || '');
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
    badge.className = 'sea-badge ice-active';
}

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

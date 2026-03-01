/* ── Survey data ────────────────────────────────────────────────── */
const SURVEYS = [
    {
        id: '2026-02-16_IMO9233387_loading',
        vessel: 'POLAR STAR', imo: '9233387', flag: 'RU',
        date: '2026-02-16', time: '08:45',
        port: 'Vanino, Russia',
        operation: 'Loading',
        cargo: 78556.093,
        declared: 78549.200,
        status: 'Complete',
        mmc_ini: 6.323, mmc_fin: 12.203,
    },
    {
        id: '2026-02-10_IMO9233387_discharging',
        vessel: 'POLAR STAR', imo: '9233387', flag: 'RU',
        date: '2026-02-10', time: '14:20',
        port: 'Nakhodka, Russia',
        operation: 'Discharging',
        cargo: 62184.500,
        declared: 62200.000,
        status: 'Complete',
        mmc_ini: 11.842, mmc_fin: 5.971,
    },
    {
        id: '2026-01-28_IMO9357432_loading',
        vessel: 'DD VIGILANT', imo: '9357432', flag: 'PA',
        date: '2026-01-28', time: '09:10',
        port: 'Murmansk, Russia',
        operation: 'Loading',
        cargo: 55321.410,
        declared: 55300.000,
        status: 'Complete',
        mmc_ini: 5.210, mmc_fin: 11.044,
    },
    {
        id: '2026-01-15_IMO9357432_discharging',
        vessel: 'DD VIGILANT', imo: '9357432', flag: 'PA',
        date: '2026-01-15', time: '11:30',
        port: 'Rotterdam, Netherlands',
        operation: 'Discharging',
        cargo: 54890.000,
        declared: 55321.410,
        status: 'Complete',
        mmc_ini: 10.998, mmc_fin: 5.100,
    },
    {
        id: '2026-01-04_IMO9741666_loading',
        vessel: 'M CONFIDANTE', imo: '9741666', flag: 'PA',
        date: '2026-01-04', time: '16:00',
        port: 'Yuzhno-Sakhalinsk, Russia',
        operation: 'Loading',
        cargo: 44102.800,
        declared: 44100.000,
        status: 'Complete',
        mmc_ini: 4.880, mmc_fin: 9.615,
    },
    {
        id: '2026-02-20_IMO9233387_loading',
        vessel: 'OCEAN CHEERS', imo: '9233380', flag: 'MH',
        date: '2026-02-20', time: '07:30',
        port: 'Vanino, Russia',
        operation: 'Loading',
        cargo: null,
        declared: 78000.000,
        status: 'In Progress',
        mmc_ini: 6.520, mmc_fin: null,
    },
    {
        id: '2026-02-24_IMO9481200_loading',
        vessel: 'BALTIC QUEEN', imo: '9481200', flag: 'CY',
        date: '2026-02-24', time: '10:15',
        port: 'St. Petersburg, Russia',
        operation: 'Loading',
        cargo: null,
        declared: 31500.000,
        status: 'In Progress',
        mmc_ini: 4.210, mmc_fin: null,
    },
    {
        id: '2026-02-25_IMO9612843_loading',
        vessel: 'GOLDEN HORIZON', imo: '9612843', flag: 'LR',
        date: '2026-02-25', time: '—',
        port: 'Ust-Luga, Russia',
        operation: 'Loading',
        cargo: null,
        declared: null,
        status: 'Draft',
        mmc_ini: null, mmc_fin: null,
    },
];

/* ── Helpers ────────────────────────────────────────────────────── */
function fmtCargo(v) {
    if (v == null) return '—';
    return v.toLocaleString('en-US', { minimumFractionDigits: 3, maximumFractionDigits: 3 });
}

function fmtDisc(cargo, declared) {
    if (cargo == null || declared == null) return { text: '—', cls: '' };
    const d = cargo - declared;
    const pct = (d / declared * 100);
    const sign = d >= 0 ? '+' : '';
    const cls = Math.abs(pct) < 0.05 ? 'ok' : d > 0 ? 'pos' : 'neg';
    return { text: sign + d.toFixed(3) + ' MT / ' + sign + pct.toFixed(3) + '%', cls };
}

function opBadge(op) {
    return op === 'Loading'
        ? '<span class="badge badge--load">Loading</span>'
        : '<span class="badge badge--disch">Discharge</span>';
}

function statusBadge(s) {
    if (s === 'Complete') return '<span class="badge badge--complete">Complete</span>';
    if (s === 'In Progress') return '<span class="badge badge--progress">In Progress</span>';
    return '<span class="badge badge--draft">Draft</span>';
}

function fmtDate(d) {
    const dt = new Date(d + 'T00:00:00');
    return dt.toLocaleDateString('en-GB', { day: '2-digit', month: 'short', year: 'numeric' });
}

/* ── Filter ─────────────────────────────────────────────────────── */
function getFiltered() {
    const q = document.getElementById('search').value.toLowerCase().trim();
    const op = document.getElementById('filter-op').value;
    const st = document.getElementById('filter-status').value;
    return SURVEYS.filter(s => {
        if (op && s.operation !== op) return false;
        if (st && s.status !== st) return false;
        if (q && ![s.vessel, s.imo, s.port].some(v => v.toLowerCase().includes(q))) return false;
        return true;
    });
}

/* ── Render table ───────────────────────────────────────────────── */
function renderTable(list) {
    const tbody = document.getElementById('tbl-body');
    const empty = document.getElementById('tbl-empty');
    tbody.innerHTML = '';
    if (!list.length) { empty.classList.add('show'); return; }
    empty.classList.remove('show');

    var pdfIcon = '<svg viewBox="0 0 24 24"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><\/svg>';
    var openIcon = '<svg viewBox="0 0 24 24"><path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"/><polyline points="15 3 21 3 21 9"/><line x1="10" y1="14" x2="21" y2="3"/><\/svg>';
    var moreIcon = '<svg viewBox="0 0 24 24"><circle cx="12" cy="5" r="1"/><circle cx="12" cy="12" r="1"/><circle cx="12" cy="19" r="1"/><\/svg>';

    list.forEach(function (s) {
        const disc = fmtDisc(s.cargo, s.declared);
        const canPDF = s.status === 'Complete';
        const tr = document.createElement('tr');

        var pdfBtn = canPDF
            ? '<button class="ia ia--pdf" title="Download PDF">' + pdfIcon + '</button>'
            : '';

        tr.innerHTML =
            '<td>' +
            '<div class="td-vessel">' +
            '<span class="td-vessel-name">' + s.vessel + '</span>' +
            '<span class="td-vessel-imo">IMO ' + s.imo + '</span>' +
            '</div>' +
            '</td>' +
            '<td>' +
            '<div class="td-date-main">' + fmtDate(s.date) + '</div>' +
            '<div class="td-date-time">' + s.time + '</div>' +
            '</td>' +
            '<td class="td-port">' + s.port + '</td>' +
            '<td>' + opBadge(s.operation) + '</td>' +
            '<td class="td-r td-cargo">' + fmtCargo(s.cargo) + '</td>' +
            '<td class="td-r td-disc ' + disc.cls + '">' + disc.text + '</td>' +
            '<td>' + statusBadge(s.status) + '</td>' +
            '<td>' +
            '<div class="row-actions">' +
            '<button class="ia ia--open" title="Open survey">' + openIcon + '</button>' +
            pdfBtn +
            '<button class="ia" title="More options">' + moreIcon + '</button>' +
            '</div>' +
            '</td>';

        tbody.appendChild(tr);
    });
}

/* ── Render cards ───────────────────────────────────────────────── */
function renderGrid(list) {
    const grid = document.getElementById('grid-view');
    grid.innerHTML = '';
    if (!list.length) return;

    list.forEach(function (s) {
        const disc = fmtDisc(s.cargo, s.declared);
        const card = document.createElement('div');
        card.className = 'survey-card';

        var discHtml = disc.text !== '—'
            ? '<span class="sc-disc ' + disc.cls + '">' + (disc.text.split('/')[1] || '').trim() + '</span>'
            : '';

        var cargoVal = s.cargo != null
            ? s.cargo.toLocaleString('en-US', { minimumFractionDigits: 3 })
            : '—';

        var cargoUnit = s.cargo != null ? '<span class="sc-cargo-unit">MT</span>' : '';

        var mmcIni = s.mmc_ini ? s.mmc_ini.toFixed(3) + ' m' : '—';
        var mmcFin = s.mmc_fin ? s.mmc_fin.toFixed(3) + ' m' : '—';
        var mmcIniCls = s.mmc_ini ? '' : 'dim';
        var mmcFinCls = s.mmc_fin ? '' : 'dim';
        var cargoLabel = s.operation === 'Loading' ? 'Cargo Loaded' : 'Cargo Discharged';

        card.innerHTML =
            '<div class="sc-header">' +
            '<div>' +
            '<div class="sc-vessel">' + s.vessel + '</div>' +
            '<div class="sc-imo">IMO ' + s.imo + '</div>' +
            '</div>' +
            '<div class="sc-badges">' +
            opBadge(s.operation) +
            statusBadge(s.status) +
            '</div>' +
            '</div>' +
            '<div class="sc-body">' +
            '<div class="sc-row"><span class="sc-lbl">Date</span><span class="sc-val">' + fmtDate(s.date) + '</span></div>' +
            '<div class="sc-row"><span class="sc-lbl">Port</span><span class="sc-val">' + s.port + '</span></div>' +
            '<div class="sc-row"><span class="sc-lbl">MMC Initial</span><span class="sc-val ' + mmcIniCls + '">' + mmcIni + '</span></div>' +
            '<div class="sc-row"><span class="sc-lbl">MMC Final</span><span class="sc-val ' + mmcFinCls + '">' + mmcFin + '</span></div>' +
            '</div>' +
            '<div class="sc-cargo">' +
            '<div>' +
            '<div class="sc-cargo-lbl">' + cargoLabel + '</div>' +
            '<div><span class="sc-cargo-val">' + cargoVal + '</span>' + cargoUnit + '</div>' +
            '</div>' +
            discHtml +
            '</div>';

        grid.appendChild(card);
    });
}

/* ── Main render ────────────────────────────────────────────────── */
function filterSurveys() {
    const list = getFiltered();
    document.getElementById('results-count').innerHTML =
        '<b>' + list.length + '</b> survey' + (list.length !== 1 ? 's' : '') + ' found';
    renderTable(list);
    renderGrid(list);
}

/* ── View toggle ────────────────────────────────────────────────── */
function setView(v) {
    const tableWrap = document.getElementById('table-view');
    const gridWrap = document.getElementById('grid-view');
    const btnTable = document.getElementById('vt-table');
    const btnGrid = document.getElementById('vt-grid');

    if (v === 'table') {
        tableWrap.style.display = '';
        gridWrap.classList.remove('active');
        btnTable.classList.add('active');
        btnGrid.classList.remove('active');
    } else {
        tableWrap.style.display = 'none';
        gridWrap.classList.add('active');
        btnTable.classList.remove('active');
        btnGrid.classList.add('active');
    }
}

/* ── Init ───────────────────────────────────────────────────────── */
filterSurveys();
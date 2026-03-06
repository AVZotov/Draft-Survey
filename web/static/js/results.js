/* ── Step collapse ──────────────────────────────────────────── */
function toggleStep(header) {
  const body = header.nextElementSibling;
  const collapsed = header.classList.toggle('collapsed');
  body.classList.toggle('hidden', collapsed);
}

/* ── Survey selector ────────────────────────────────────────── */
function updateSelector() {
  // TODO: HTMX request to backend with selected draft IDs
  console.log('ini:', document.getElementById('sel-ini').value,
              'fin:', document.getElementById('sel-fin').value);
}

/* ── Print menu ─────────────────────────────────────────────── */
function togglePrintMenu() {
  document.getElementById('print-menu').classList.toggle('open');
}

function printReport(type) {
  document.getElementById('print-menu').classList.remove('open');
  // TODO: request to backend for PDF generation
  alert('Print: ' + type + '\n(Backend PDF generation)');
}

document.addEventListener('click', function(e) {
  const wrap = document.querySelector('.print-wrap');
  if (wrap && !wrap.contains(e.target)) {
    document.getElementById('print-menu').classList.remove('open');
  }
});

/* ── Edit panel ─────────────────────────────────────────────── */
function openEditPanel() {
  document.getElementById('panel-overlay').classList.add('open');
  document.body.style.overflow = 'hidden';
}

function closeEditPanel() {
  document.getElementById('panel-overlay').classList.remove('open');
  document.body.style.overflow = '';
}

function switchTab(id, btn) {
  document.querySelectorAll('.panel-tab-content').forEach(t => t.classList.remove('active'));
  document.querySelectorAll('.panel-tab').forEach(t => { t.classList.remove('active', 'fin'); });
  document.getElementById('tab-' + id).classList.add('active');
  btn.classList.add('active');
  if (id === 'fin') btn.classList.add('fin');
}

function resetPanel() {
  // TODO: reload data from backend
  if (confirm('Reset all changes to last saved values?')) {
    closeEditPanel();
  }
}

function submitPanel() {
  // TODO: HTMX POST to backend for recalculation
  closeEditPanel();
  alert('Submitted to backend for recalculation');
}

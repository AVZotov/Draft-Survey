function toggleSeaGroup(group, prefix) {
    const other = group === 'wave' ? 'ice' : 'wave';

    // switch on selected
    document.getElementById(prefix + '-' + group + '-group').classList.remove('sea-selects--disabled');
    document.getElementById(group === 'wave' ? prefix + '-sea-condition' : prefix + '-ice-condition').disabled = false;

    // switch on other
    document.getElementById(prefix + '-' + other + '-group').classList.add('sea-selects--disabled');
    document.getElementById(other === 'wave' ? prefix + '-sea-condition' : prefix + '-ice-condition').disabled = true;
}

// validate elements per page changes form
document.getElementById('pagesize').onchange = function() {
    console.debug("pagesize onchange event");
    document.getElementById('pagesize-form').submit();
};

// validate current form
document.getElementById('page').onchange = function() {
    console.debug("page onchange event");
    document.getElementById('page-form').submit();
};
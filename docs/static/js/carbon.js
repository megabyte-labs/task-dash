(function () {
  /**
   *
   */
  function attachAd() {
    const element = document.createElement('script');
    element.setAttribute('type', 'text/javascript');
    element.setAttribute('id', '_carbonads_js');
    element.setAttribute(
      'src',
      '//cdn.carbonads.com/carbon.js?serve=CESI65QJ&placement=taskfiledev'
    );
    element.setAttribute('async', 'async');

    const wrapper = document.querySelector('#sidebar-ads');
    wrapper.innerHTML = '';
    wrapper.append(element);
  }

  setTimeout(function () {
    attachAd();

    window.addEventListener('popstate', function () {
      attachAd();
    });
  }, 1000);
})();

import { Component, Vue } from 'vue-property-decorator';
import sanitizeHtml from 'sanitize-html';
import { deviceType, DeviceType } from '@/models/types/deviceType';
import store from '@/store';
import { TYPES } from '@/store/mutation-types';

@Component
export default class UtilMixin extends Vue {
  deepCopyUtil<T>(obj: T): T {
    return JSON.parse(JSON.stringify(obj));
  }
  sanitizeUtil = (value: string): string => {
    return value ? sanitizeHtml(value, { allowedTags: ['br', 'a'] }) : '';
  };
  scrollUtil(target?: string) {
    if (target) {
      store.dispatch(TYPES.SCROLL_TARGET, target);
    }
    if (store.state.scrollTarget) {
      this.$vuetify.goTo(store.state.scrollTarget, {
        duration: 1000,
        easing: 'easeInOutQuint',
      });
    }
    store.dispatch(TYPES.SCROLL_TARGET, '');
  }
  scrollVisibilityUtil(element: HTMLElement | null) {
    const elementTop = element
      ? element.getBoundingClientRect().top
      : window.innerHeight;
    const elementHeight = element
      ? element.getBoundingClientRect().height
      : window.innerHeight;
    return elementTop + elementHeight * 0.2 < window.innerHeight;
  }
  visibilityUtil(isVisible: boolean) {
    return isVisible ? 'visibility: visible;' : 'visibility: hidden;';
  }
  MdSmXsUtil(md: number | string, sm: number | string, xs: number | string) {
    return deviceType === DeviceType.DESKTOP
      ? md
      : deviceType === DeviceType.TABLET
      ? sm
      : xs;
  }
  fontSizeUtil(mdFontSize: number, smFontSize: number, xsFontSize: number) {
    return `font-size: ${this.MdSmXsUtil(
      mdFontSize,
      smFontSize,
      xsFontSize
    )}px;`;
  }
  maxWidthUtil(mdMaxWidth: string, smMaxWidth: string, xsMaxWidth: string) {
    return `max-width: ${this.MdSmXsUtil(mdMaxWidth, smMaxWidth, xsMaxWidth)};`;
  }
  openBrowserTabUtil(url: string) {
    window.open(url);
  }
  linkUtil(label: string, href: string) {
    return `<a href="${href}" target="_blank" rel="noopener noreferrer">${label}</a>`;
  }
  sleep = (msec: number) => new Promise((resolve) => setTimeout(resolve, msec));
}

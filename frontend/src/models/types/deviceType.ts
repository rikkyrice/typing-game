export enum DeviceType {
  DESKTOP = 'desktop',
  TABLET = 'tablet',
  MOBILE = 'mobile',
}

export const deviceType = /iPad/.test(navigator.userAgent)
  ? DeviceType.TABLET
  : /Mobile|iP(hone|od)|Android|BlackBerry|IEMobile|Silk/.test(
      navigator.userAgent
    )
  ? DeviceType.MOBILE
  : DeviceType.DESKTOP;

/**
 * Admin API barrel export
 * Centralized exports for all admin API modules
 */

import dashboardAPI from './dashboard'
import usersAPI from './users'
import groupsAPI from './groups'
import accountsAPI from './accounts'
import proxiesAPI from './proxies'
import settingsAPI from './settings'
import systemAPI from './system'
import subscriptionsAPI from './subscriptions'
import usageAPI from './usage'
import geminiAPI from './gemini'
import antigravityAPI from './antigravity'
import apiKeysAPI from './apiKeys'
import channelsAPI from './channels'

/**
 * Unified admin API object for convenient access
 */
export const adminAPI = {
  dashboard: dashboardAPI,
  users: usersAPI,
  groups: groupsAPI,
  accounts: accountsAPI,
  proxies: proxiesAPI,
  settings: settingsAPI,
  system: systemAPI,
  subscriptions: subscriptionsAPI,
  usage: usageAPI,
  gemini: geminiAPI,
  antigravity: antigravityAPI,
  apiKeys: apiKeysAPI,
  channels: channelsAPI
}

export {
  dashboardAPI,
  usersAPI,
  groupsAPI,
  accountsAPI,
  proxiesAPI,
  settingsAPI,
  systemAPI,
  subscriptionsAPI,
  usageAPI,
  geminiAPI,
  antigravityAPI,
  apiKeysAPI,
  channelsAPI
}

export default adminAPI

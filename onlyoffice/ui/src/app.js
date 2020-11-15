const appInfo = {
  name: 'ocis-onlyoffice',
  id: 'onlyoffice',
  icon: 'x-office-document',
  extensions: [
    {
      extension: 'docx',
      handler: function(config, filePath, fileId) {
        window.open(
          `${config.server}/apps/onlyoffice/${fileId}?filePath=${encodeURIComponent(filePath)}`,
          '_blank'
        )
      },
      newFileMenu: {
        menuTitle($gettext) {
          return $gettext('New Onlyoffice document')
        }
      }
    },
    {
      extension: 'xlxs',
      handler: function(config, filePath, fileId) {
        window.open(
          `${config.server}/apps/onlyoffice/${fileId}?filePath=${encodeURIComponent(filePath)}`,
          '_blank'
        )
      },
      newFileMenu: {
        menuTitle($gettext) {
          return $gettext('New Onlyoffice spreadsheet')
        }
      }
    },
    {
      extension: 'pptx',
      handler: function(config, filePath, fileId) {
        window.open(
          `${config.server}/apps/onlyoffice/${fileId}?filePath=${encodeURIComponent(filePath)}`,
          '_blank'
        )
      },
      newFileMenu: {
        menuTitle($gettext) {
          return $gettext('New Onlyoffice presentation')
        }
      }
    }
  ]
}

export default define({
  appInfo
})


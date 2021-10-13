import { useState } from 'react'
import { createTheme, ThemeProvider } from '@mui/material/styles'
import { blue } from '@mui/material/colors'
import { SnackbarProvider } from 'notistack'

import Tabs from '@mui/material/Tabs'
import Tab from '@mui/material/Tab'
import Box from '@mui/material/Box'

import Servers from './components/servers'
import Nodes from './components/nodes'
import Rules from './components/rules'
import Summary from './components/summary'

const myTheme = createTheme({
  palette: {
    primary: {
      main: blue[800],
    },
  },
})

export const DefaultMsgOption: any = {
  autoHideDuration: 1000,
  anchorOrigin: {
    vertical: 'top',
    horizontal: 'center',
  }
}

export default function App() {
  const [tabIndex, setTabIndex] = useState(0)
  const switchTab = (event: React.SyntheticEvent, newValue: number) => {
    setTabIndex(newValue)
  }

  const tabList: any = [
    { label: "总览", comp: <Summary /> },
    { label: "节点", comp: <Nodes /> },
    { label: "服务器", comp: <Servers /> },
    { label: "规则", comp: <Rules /> },
  ]

  return (
    <ThemeProvider theme={myTheme}>
      <SnackbarProvider maxSnack={3}>
        <Box sx={{ width: '100%' }}>
          <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
            <Tabs value={tabIndex} onChange={switchTab} centered>
              {tabList.map((data: any, index: number) => {
                return (
                  <Tab
                    key={"simple-tab-" + index}
                    id={"simple-tab-" + index}
                    label={data.label}
                    aria-controls={"simple-tabpanel-" + index}
                  />
                )
              })}
            </Tabs>
          </Box>
          {tabList[tabIndex].comp}
        </Box>
      </SnackbarProvider>
    </ThemeProvider>
  )
}

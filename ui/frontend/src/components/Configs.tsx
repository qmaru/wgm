import { useState, useEffect, useCallback } from 'react'

import Container from '@mui/material/Container'
import Box from '@mui/material/Box'
import Paper from '@mui/material/Paper'
import Stepper from '@mui/material/Stepper'
import Step from '@mui/material/Step'
import StepLabel from '@mui/material/StepLabel'
import Button from '@mui/material/Button'
import Typography from '@mui/material/Typography'
import Radio from '@mui/material/Radio'
import RadioGroup from '@mui/material/RadioGroup'
import FormGroup from '@mui/material/FormGroup'
import FormControlLabel from '@mui/material/FormControlLabel'
import FormControl from '@mui/material/FormControl'
import FormLabel from '@mui/material/FormLabel'
import Checkbox from '@mui/material/Checkbox'
import Stack from '@mui/material/Stack'
import Select from '@mui/material/Select'
import Chip from '@mui/material/Chip'
import Dialog from '@mui/material/Dialog'
import DialogContent from '@mui/material/DialogContent'
import MenuItem from '@mui/material/MenuItem'
import OutlinedInput from '@mui/material/OutlinedInput'
import InputLabel from '@mui/material/InputLabel'
import Divider from '@mui/material/Divider'
import { useTheme } from '@mui/material/styles'

import { useSnackbar } from 'notistack'
import QRCode from "react-qr-code"

import { ClipboardText, DataListAPI } from "../../wailsjs/go/backend/App"

const steps = [
  "选择 Interface",
  "选择 Peers",
  "确认配置"
]

const InterfaceOption = (peer: any) => {
  const username = peer.username
  const private_addr = peer.private_addr
  const public_addr = peer.public_addr
  const port = peer.port
  let public_info = ""
  if (public_addr !== "") {
    public_info = `${public_addr}:${port}`
  }
  return (
    <Stack
      divider={<Divider orientation="vertical" flexItem />}
      direction="row"
      spacing={1}
    >
      <Typography color={public_addr === "" ? "default" : "secondary"}>
        {username}
      </Typography>
      <Typography color={public_addr === "" ? "default" : "secondary"}>
        {private_addr}
      </Typography>
      {public_info === "" ? null : <Typography color={public_addr === "" ? "default" : "secondary"}>{public_info}</Typography>}
    </Stack>
  )
}


export default function Configs() {
  const theme = useTheme()
  const { enqueueSnackbar } = useSnackbar()

  const [activeStep, setActiveStep] = useState(0)

  const [qrOpen, setQROpen] = useState<boolean>(false)
  const [qrData, setQrData] = useState<string>("")
  const [copyOpen, setCopyOpen] = useState<boolean>(false)
  const [copyData, setCopyData] = useState<string>("")

  const [peerData, setPeerData] = useState<any>([])
  const [routeData, setRouteData] = useState<any>([])

  const [interfaceNode, setInterfaceNode] = useState<string>("")
  const [peerNodes, setPeerNodes] = useState<any>([])
  const [routeNodes, setRouteNodes] = useState<any>([])
  const [interfaceConfig, setInterfaceConfig] = useState<any>([])
  const [peerConfigs, setPeerConfigs] = useState<any>([])

  const QRClose = () => {
    setQROpen(false)
  }

  const QROpen = () => {
    setQROpen(true)
  }

  const InterfaceNodeChange = (event: any) => {
    setInterfaceNode(event.target.value)
  }

  const PeerNodesChange = (event: any) => {
    setPeerNodes({ ...peerNodes, [event.target.name]: event.target.checked })
  }

  const RouteNodesChange = (event: any, selectId: number) => {
    const value = event.target.value
    setRouteNodes((prevValues: any) => ({
      ...prevValues,
      [selectId]: value,
    }))
  }

  const StepNext = () => {
    if (activeStep === 0 && interfaceNode === "") {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请选择 Interface",
        window.messageDefault
      )
      return false
    }
    setActiveStep((prevActiveStep) => prevActiveStep + 1)
  }

  const StepBack = () => {
    if (activeStep === 1) {
      setInterfaceNode("")
    } else if (activeStep === 2) {
      setPeerNodes([])
      setRouteNodes([])
    }
    setActiveStep((prevActiveStep) => prevActiveStep - 1)
  }

  const StepReset = () => {
    setActiveStep(0)
    setInterfaceNode("")
    setPeerNodes([])
    setRouteNodes([])
    setCopyOpen(false)
  }

  const ConfigOut = () => {
    if (activeStep === 1 && peerNodes.length === 0) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请选择 Peers",
        window.messageDefault
      )
      return false
    }

    setActiveStep((prevActiveStep) => prevActiveStep + 1)
    let configInterface: any = {}
    let configPeers: any = []
    for (let i in peerData) {
      let pData = peerData[i]
      if (pData.id.toString() === interfaceNode) {
        configInterface = pData
      }
    }

    let peerNodesTrue: any = []
    for (let i in peerNodes) {
      if (peerNodes[i]) {
        peerNodesTrue.push(i)
      }
    }

    for (let i in peerData) {
      let pData = peerData[i]
      if (peerNodesTrue.includes(pData.id.toString())) {
        configPeers.push(pData)
      }
    }

    let interface_port = ""
    let interface_mtu = ""
    let interface_dns = ""

    if (configInterface.port !== 0) {
      interface_port = `ListenPort = ${configInterface.port}`
    }

    if (configInterface.mtu !== 0) {
      interface_mtu = `MTU = ${configInterface.mtu}`
    }

    if (configInterface.dns !== "") {
      interface_dns = `DNS = ${configInterface.dns}`
    }

    const interface_section = "[Interface]"
    const interface_remark = `# ${configInterface.username}`
    const interface_private_key = `PrivateKey = ${configInterface.prikey}`
    const interface_address = `Address = ${configInterface.private_addr}/24`

    const interface_config = [
      interface_section,
      interface_remark,
      interface_private_key,
      interface_address,
      interface_port,
      interface_mtu,
      interface_dns
    ]

    let peer_configs = []
    for (let i in configPeers) {
      let peer_tmp = []
      const pData = configPeers[i]
      const peer_id = pData.id.toString()
      const allowed_ips = pData.allowed_ips
      let extar_ips = [allowed_ips]
      if (Object.keys(routeNodes).includes(peer_id)) {
        extar_ips.push(...routeNodes[peer_id])
      }
      const peer_section = "[Peer]"
      const peer_remark = `# ${pData.username}`
      const peer_public_key = `PublicKey = ${pData.pubkey}`
      const peer_allowed_ips = `AllowedIPs = ${extar_ips.toString()}`
      let peer_endpoint = ""
      let peer_keepalive = ""
      if (pData.public_addr !== "") {
        peer_endpoint = `Endpoint = ${pData.public_addr}:${pData.port}`
      }
      if (pData.port !== 0) {
        peer_keepalive = `PersistentKeepalive = ${pData.keepalive}`
      }
      peer_tmp.push(peer_section)
      peer_tmp.push(peer_remark)
      peer_tmp.push(peer_public_key)
      peer_tmp.push(peer_allowed_ips)
      peer_tmp.push(peer_endpoint)
      peer_tmp.push(peer_keepalive)
      const peer_tmp_f = peer_tmp.filter((item: any) => item !== "")
      peer_configs.push(peer_tmp_f)
    }

    const interface_config_f = interface_config.filter((item: any) => item !== "")

    const interface_config_str = interface_config_f.join("\n")
    const peer_config_str = peer_configs.map((config: any) => config.join("\n")).join("\n\n")

    const configDataStr = interface_config_str + "\n\n" + peer_config_str + "\n\n"
    setCopyData(configDataStr)
    setQrData(configDataStr)

    setInterfaceConfig(interface_config_f)
    setPeerConfigs(peer_configs)
  }

  const PeerDataList = useCallback(() => {
    DataListAPI()
      .then(response => {
        let status = response.status
        if (status === 1) {
          let data = response.data
          data.routes.sort((i: any, j: any) => i.cidr.localeCompare(j.cidr))
          data.peers.sort((i: any, j: any) => {
            const ipSegments1 = i.private_addr.split('.').map((segment: any) => parseInt(segment))
            const ipSegments2 = j.private_addr.split('.').map((segment: any) => parseInt(segment))
            for (let i = 0; i < 4; i++) {
              if (ipSegments1[i] < ipSegments2[i]) {
                return -1
              } else if (ipSegments1[i] > ipSegments2[i]) {
                return 1
              }
            }
            return 0
          })
          setRouteData(data.routes)
          setPeerData(data.peers)
        }
      })
      .catch(
        () => {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            "配置接口请求失败",
            window.messageDefault
          )
        }
      )
  }, [enqueueSnackbar])

  const CopyConfig = () => {
    ClipboardText(copyData)
      .then((res: boolean) => {
        if (res) {
          setCopyOpen(true)
        }
      })
  }

  useEffect(() => {
    PeerDataList()
  }, [PeerDataList])

  return (
    <Container key={"Peers-Main"} disableGutters maxWidth={false}>

      <Container key={"Peers-Stepper"}
        sx={{ p: 4 }}
      >
        <Stepper activeStep={activeStep}>
          {steps.map((step: any, index: number) => {
            return (
              <Step key={"step" + index} >
                <StepLabel>{step}</StepLabel>
              </Step>
            )
          })}
        </Stepper>
      </Container>

      <Container key={"Peers-Step-Main"} sx={{ pb: 4 }}>
        {activeStep === steps.length ?
          <Box sx={{ display: 'flex', flexDirection: 'row' }}>
            <Button
              color="error"
              variant="contained"
              onClick={StepReset}>
              重置
            </Button>
            <Box sx={{ flex: '1 1 auto' }} />

            <Button
              variant="contained"
              color='success'
              onClick={() => CopyConfig()}
            >
              {copyOpen ? "已复制" : "复制"}
            </Button>
          </Box>
          :
          <Box sx={{ display: 'flex', flexDirection: 'row' }}>
            <Button
              variant="contained"
              disabled={activeStep === 0}
              onClick={StepBack}
            >
              上一步
            </Button>
            <Box sx={{ flex: '1 1 auto' }} />
            <Button
              variant="contained"
              onClick={activeStep === 1 ? ConfigOut : StepNext}
            >
              {activeStep === steps.length - 1 ? '完成' : '下一步'}
            </Button>
          </Box>
        }
        <Box sx={{ paddingTop: 3 }}>
          {activeStep === 0 ?
            <FormControl>
              <FormLabel sx={{ paddingBottom: 1 }}>选择<b>本地</b>接口 Interface</FormLabel>
              <RadioGroup
                value={interfaceNode}
                onChange={InterfaceNodeChange}
              >
                {peerData.map((peer: any, index: number) => {
                  return (
                    <FormControlLabel
                      key={"interface" + index}
                      sx={{ pt: 1 }}
                      value={peer.id}
                      control={<Radio color={peer.public_addr === "" ? "primary" : "secondary"} />}
                      label={InterfaceOption(peer)}
                    />
                  )
                })}
              </RadioGroup>
            </FormControl>
            : activeStep === 1 ?
              <FormGroup>
                <FormLabel>选择<b>对端</b> Peer</FormLabel>
                <Stack>
                  {peerData.map((peer: any, index: number) => {
                    return (
                      <Box key={"peer" + index}>
                        <Stack
                          direction="row"
                          justifyContent="space-between"
                          alignItems="center"
                        >
                          <Box>
                            <FormControlLabel
                              label={InterfaceOption(peer)}
                              control={
                                <Checkbox
                                  name={peer.id.toString()}
                                  color={peer.public_addr === "" ? "primary" : "secondary"}
                                  onChange={PeerNodesChange}
                                />
                              }
                            />
                          </Box>

                          <Box sx={{ paddingBottom: 2, paddingTop: 2 }}>
                            <FormControl sx={{ width: 400 }}>
                              <InputLabel>路由规则</InputLabel>
                              <Select
                                multiple
                                value={routeNodes[peer.id.toString()] || []}
                                onChange={(event) => RouteNodesChange(event, peer.id.toString())}
                                input={<OutlinedInput label="Chip" />}
                                renderValue={(selected) => (
                                  <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                                    {selected.map((value: any) => (
                                      <Chip key={value} label={value} />
                                    ))}
                                  </Box>
                                )}
                              >
                                {routeData.map((route: any) => (
                                  <MenuItem
                                    key={"route" + route.id}
                                    value={route.cidr}
                                  >
                                    {route.cidr}
                                  </MenuItem>
                                ))}
                              </Select>
                            </FormControl>
                          </Box>
                        </Stack>
                        <Divider />
                      </Box>
                    )
                  })}
                </Stack>
              </FormGroup>
              :
              <Box>
                <Box sx={{ paddingBottom: 4 }}>
                  <Button
                    variant="contained"
                    color="secondary"
                    onClick={() => QROpen()}
                    sx={{ color: "white" }}
                    disabled={activeStep !== steps.length}
                  >
                    显示二维码
                  </Button>
                </Box>
                <Paper variant="outlined" sx={{ borderBottomLeftRadius: 8, borderBottomRightRadius: 8 }}>
                  <Box
                    sx={{
                      borderTopLeftRadius: 8,
                      borderTopRightRadius: 8,
                      backgroundColor: theme.palette.primary.main,
                      p: 0.6
                    }}
                  />
                  <Box sx={{ p: 2 }}>
                    {interfaceConfig.map((inter_c: any, index: number) => {
                      return (
                        <Typography key={"intercfg" + index}>{inter_c}</Typography>
                      )
                    })}
                    <Box m={2} />
                    {peerConfigs.map((peer_config: any, index: number) => {
                      return (
                        <Box key={"peerscfg" + index}>
                          {peer_config.map((peer: any, index: number) => {
                            return (
                              <Typography key={"peercfg" + index}>{peer}</Typography>
                            )
                          })}
                          <Box m={2} />
                        </Box>
                      )
                    })}
                  </Box>
                </Paper>
              </Box>
          }
        </Box>
      </Container>

      <Dialog open={qrOpen} onClose={QRClose} >
        <DialogContent sx={{ m: 0 }}>
          <QRCode value={qrData} size={256} fgColor="#000000" />
        </DialogContent>
      </Dialog >

    </Container>
  )
}

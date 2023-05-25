import { useState, useEffect, useCallback } from 'react'

import Container from '@mui/material/Container'
import Box from '@mui/material/Box'
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
import MenuItem from '@mui/material/MenuItem'
import OutlinedInput from '@mui/material/OutlinedInput'
import InputLabel from '@mui/material/InputLabel'

import { useSnackbar } from 'notistack'

export default function Peers() {
  const { enqueueSnackbar } = useSnackbar()

  const [activeStep, setActiveStep] = useState(0)
  const [peerData, setPeerData] = useState<any>([])
  const [routeData, setRouteData] = useState<any>([])
  const [interfaceNode, setInterfaceNode] = useState<string>("")
  const [peerNodes, setPeerNodes] = useState<any>([])
  const [routeNodes, setRouteNodes] = useState<any>([])

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

  const steps = [
    "选择 Interface",
    "选择 Peers",
    "确认配置"
  ]

  const StepNext = () => {
    if (activeStep === 0 && interfaceNode === "") {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请选择 Interface",
        window.messageDefault
      )
      return
    } else if (activeStep === 1 && peerNodes.length === 0) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请选择 Peers",
        window.messageDefault
      )
      return
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
  }

  const StepCopy = () => {
    alert("复制了一些配置")
  }

  const InterfaceOption = (peer: any) => {
    const username = peer.username
    const private_addr = peer.private_addr
    const public_addr = peer.public_addr
    const port = peer.port
    let public_info = ""
    if (public_addr !== "") {
      public_info = ` | ${public_addr}:${port}`
    }
    const option = `${username} | ${private_addr}${public_info}`
    return option
  }

  const ConfigOut = () => {
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
      interface_port = `ListenPort =  ${configInterface.port}`
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
      peer_configs.push(peer_tmp)
    }

    return (
      <Box>
        {interface_config.map((inter_c: any, index: number) => {
          return (
            <Typography key={"inter" + index}>{inter_c}</Typography>
          )
        })}
        <Box m={2} />
        {peer_configs.map((peer_config: any, index: number) => {
          return (
            <Box key={"peer_g" + index}>
              {peer_config.map((peer: any, index: number) => {
                return (
                  <Typography key={"peer" + index}>{peer}</Typography>
                )
              })}
              <Box m={2} />
            </Box>
          )
        })}
      </Box>
    )
  }

  const PeerDataList = useCallback(() => {
    const url = `${window.api}/data`
    fetch(url, {
      method: "GET",
    }).then(res => res.json())
      .then(response => {
        let status = response.status
        if (status === 1) {
          let data = response.data
          setRouteData(data.routes)
          setPeerData(data.peers)
        }
      })
      .catch(
        () => {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            "用户数据载入失败",
            window.messageDefault
          )
        }
      )
  }, [enqueueSnackbar])

  useEffect(() => {
    PeerDataList()
  }, [PeerDataList])

  return (
    <Container key={"Peers-Main"}>

      <Container key={"Peers-Stepper"}
        sx={{ padding: 4 }}
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

      <Container key={"Peers-Step-Main"}>
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
              color="success"
              variant="contained"
              onClick={StepCopy}>
              复制
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
              onClick={StepNext}
            >
              {activeStep === steps.length - 1 ? '完成' : '下一步'}
            </Button>
          </Box>
        }
        <Box sx={{ padding: 4 }}>
          {activeStep === 0 ?
            <FormControl>
              <FormLabel>选择一个接口</FormLabel>
              <RadioGroup
                value={interfaceNode}
                onChange={InterfaceNodeChange}
              >
                {peerData.map((peer: any, index: number) => {
                  return (
                    <FormControlLabel
                      key={"interface" + index}
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
                <FormLabel>选择节点</FormLabel>
                <Stack>
                  {peerData.map((peer: any, index: number) => {
                    return (
                      <Stack key={"peer" + index}>
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
                          <FormControl sx={{ width: "50%" }}>
                            <InputLabel>路由</InputLabel>
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
                    )
                  })}
                </Stack>
              </FormGroup>
              :
              <ConfigOut />
          }
        </Box>
      </Container>
    </Container>
  )
}

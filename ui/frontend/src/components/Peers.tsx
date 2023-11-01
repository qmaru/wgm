import { useState, useEffect, useCallback } from 'react'

import Container from '@mui/material/Container'
import Button from '@mui/material/Button'
import TextField from '@mui/material/TextField'
import Stack from '@mui/material/Stack'
import Typography from '@mui/material/Typography'
import Dialog from '@mui/material/Dialog'
import DialogActions from '@mui/material/DialogActions'
import DialogContent from '@mui/material/DialogContent'
import DialogTitle from '@mui/material/DialogTitle'
import InputLabel from '@mui/material/InputLabel'
import MenuItem from '@mui/material/MenuItem'
import FormControl from '@mui/material/FormControl'
import Tooltip from '@mui/material/Tooltip'
import Chip from '@mui/material/Chip'
import Select from '@mui/material/Select'
import { useTheme } from '@mui/material/styles'

import { useSnackbar } from 'notistack'

import { UserListAPI, PeerListAPI, PeerAddAPI, PeerUpdateAPI, PeerDeleteAPI } from "../../wailsjs/go/backend/App"
import { MyCard } from './common'


export default function Peers() {
  const theme = useTheme()

  const { enqueueSnackbar } = useSnackbar()
  const [manualRender, setManualRender] = useState<boolean>(false)

  const [userData, setUserData] = useState<any>([])
  const [peerData, setPeerData] = useState<any>([])

  const [peerAddOpen, setPeerAddOpen] = useState<boolean>(false)
  const [peerAddUser, setPeerAddUser] = useState<string>("")
  const [peerAddPublic, setPeerAddPublic] = useState<string>("")
  const [peerAddPrivate, setPeerAddPrivate] = useState<string>("")
  const [peerAddPort, setPeerAddPort] = useState<number>(0)
  const [peerAddAllowedIP, setPeerAddAllowedIP] = useState<string>("")
  const [peerAddMtu, setPeerAddMtu] = useState<number>(0)
  const [peerAddDns, setPeerAddDns] = useState<string>("")
  const [peerAddKeepalive, setPeerAddKeepalive] = useState<number>(0)

  const [peerUpdateOpen, setPeerUpdateOpen] = useState<boolean>(false)
  const [peerUpdateID, setPeerUpdateID] = useState<number>(0)
  const [peerUpdateUser, setPeerUpdateUser] = useState<string>("")
  const [peerUpdatePublic, setPeerUpdatePublic] = useState<string>("")
  const [peerUpdatePrivate, setPeerUpdatePrivate] = useState<string>("")
  const [peerUpdatePort, setPeerUpdatePort] = useState<number>(0)
  const [peerUpdateAllowedIP, setPeerUpdateAllowedIP] = useState<string>("")
  const [peerUpdateMtu, setPeerUpdateMtu] = useState<number>(0)
  const [peerUpdateDns, setPeerUpdateDns] = useState<string>("")
  const [peerUpdateKeepalive, setPeerUpdateKeepalive] = useState<number>(0)

  const [peerDeleteOpen, setPeerDeleteOpen] = useState<boolean>(false)
  const [peerDeleteID, setPeerDeleteID] = useState<number>(0)


  const PeerAddOpen = () => {
    setPeerAddUser("")
    setPeerAddPublic("")
    setPeerAddPrivate("")
    setPeerAddPort(0)
    setPeerAddAllowedIP("")
    setPeerAddMtu(0)
    setPeerAddDns("")
    setPeerAddKeepalive(0)
    setPeerAddOpen(true)
  }

  const PeerAddClose = () => {
    setPeerAddOpen(false)
  }

  const PeerUpdateOpen = (peer_data: any) => {
    setPeerUpdateID(peer_data.id)
    setPeerUpdateUser(peer_data.username)
    setPeerUpdatePublic(peer_data.public_addr)
    setPeerUpdatePrivate(peer_data.private_addr)
    setPeerUpdatePort(peer_data.port)
    setPeerUpdateAllowedIP(peer_data.allowed_ips)
    setPeerUpdateMtu(peer_data.mtu)
    setPeerUpdateDns(peer_data.dns)
    setPeerUpdateKeepalive(peer_data.keepalive)
    setPeerUpdateOpen(true)
  }

  const PeerUpdateClose = () => {
    setPeerUpdateOpen(false)
  }

  const PeerDeleteOpen = (peer_data: any) => {
    setPeerDeleteID(peer_data.id)
    setPeerDeleteOpen(true)
  }

  const PeerDeleteClose = () => {
    setPeerDeleteOpen(false)
  }

  const PeerAddUserChange = (event: any) => {
    setPeerAddUser(event.target.value)
  }

  const PeerAddPublicChange = (event: any) => {
    setPeerAddPublic(event.target.value)
  }

  const PeerAddPrivateChange = (event: any) => {
    setPeerAddPrivate(event.target.value)
  }

  const PeerAddPortChange = (event: any) => {
    setPeerAddPort(event.target.value)
  }

  const PeerAddAllowedIPChange = (event: any) => {
    setPeerAddAllowedIP(event.target.value)
  }

  const PeerAddMtuChange = (event: any) => {
    setPeerAddMtu(event.target.value)
  }

  const PeerAddDnsChange = (event: any) => {
    setPeerAddDns(event.target.value)
  }

  const PeerAddKeepaliveChange = (event: any) => {
    setPeerAddKeepalive(event.target.value)
  }

  const PeerUpdatePublicChange = (event: any) => {
    setPeerUpdatePublic(event.target.value)
  }

  const PeerUpdatePrivateChange = (event: any) => {
    setPeerUpdatePrivate(event.target.value)
  }

  const PeerUpdatePortChange = (event: any) => {
    setPeerUpdatePort(event.target.value)
  }

  const PeerUpdateAllowedIPChange = (event: any) => {
    setPeerUpdateAllowedIP(event.target.value)
  }

  const PeerUpdateMtuChange = (event: any) => {
    setPeerUpdateMtu(event.target.value)
  }

  const PeerUpdateDnsChange = (event: any) => {
    setPeerUpdateDns(event.target.value)
  }

  const PeerUpdateKeepaliveChange = (event: any) => {
    setPeerUpdateKeepalive(event.target.value)
  }

  const PeerAdd = () => {
    if (peerAddUser === "") {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请选择用户",
        window.messageDefault
      )
      return false
    }

    var regex = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/
    if (peerAddPrivate === "") {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入内网地址",
        window.messageDefault
      )
      return false
    } else if (!regex.test(peerAddPrivate)) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入正确的IP地址",
        window.messageDefault
      )
      return false
    }

    if (peerAddPublic !== "" && peerAddPort === 0) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入公网地址和端口",
        window.messageDefault
      )
      return false
    }

    if (peerAddPublic === "" && peerAddPort !== 0) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入公网地址和端口",
        window.messageDefault
      )
      return false
    }

    if (peerAddPort > 65535 || peerAddPort < 0) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入正确的端口",
        window.messageDefault
      )
      return false
    }

    if (peerAddMtu !== 0) {
      if (peerAddMtu > 1518 || peerAddMtu < 64) {
        window.messageDefault.variant = "warning"
        enqueueSnackbar(
          "请输入正确的MTU",
          window.messageDefault
        )
        return false
      }
    }

    if (peerAddKeepalive < 0) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入正确的Keepalive",
        window.messageDefault
      )
      return false
    }

    let body: any = {
      "user_id": peerAddUser,
      "public_addr": peerAddPublic,
      "private_addr": peerAddPrivate,
      "port": Number(peerAddPort),
      "allowed_ips": peerAddAllowedIP,
      "mtu": Number(peerAddMtu),
      "dns": peerAddDns,
      "keepalive": Number(peerAddKeepalive)
    }


    PeerAddAPI(body)
      .then(response => {
        let status = response.status
        if (status === 1) {
          setManualRender(!manualRender)
          window.messageDefault.variant = "success"
          enqueueSnackbar(
            "节点添加成功",
            window.messageDefault
          )
          setPeerAddOpen(false)
        } else {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            response.message,
            window.messageDefault
          )
        }
      })
      .catch(
        () => {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            "节点数据增加失败",
            window.messageDefault
          )
        }
      )
  }

  const PeerUpdate = () => {
    var regex = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/
    if (peerUpdatePrivate === "") {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入内网地址",
        window.messageDefault
      )
      return false
    } else if (!regex.test(peerUpdatePrivate)) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入正确的IP地址",
        window.messageDefault
      )
      return false
    }

    if (peerUpdatePublic !== "" && peerUpdatePort === 0) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入公网地址和端口",
        window.messageDefault
      )
      return false
    }

    if (peerUpdatePublic === "" && peerUpdatePort !== 0) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入公网地址和端口",
        window.messageDefault
      )
      return false
    }

    if (peerUpdatePort > 65535 || peerUpdatePort < 0) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入正确的端口",
        window.messageDefault
      )
      return false
    }

    if (peerUpdateMtu !== 0) {
      if (peerUpdateMtu > 1518 || peerUpdateMtu < 64) {
        window.messageDefault.variant = "warning"
        enqueueSnackbar(
          "请输入正确的MTU",
          window.messageDefault
        )
        return false
      }
    }

    if (peerUpdateKeepalive < 0) {
      window.messageDefault.variant = "warning"
      enqueueSnackbar(
        "请输入正确的Keepalive",
        window.messageDefault
      )
      return false
    }

    let body: any = {
      "public_addr": peerUpdatePublic,
      "private_addr": peerUpdatePrivate,
      "port": Number(peerUpdatePort),
      "allowed_ips": peerUpdateAllowedIP,
      "mtu": Number(peerUpdateMtu),
      "dns": peerUpdateDns,
      "keepalive": Number(peerUpdateKeepalive)
    }

    PeerUpdateAPI(String(peerUpdateID), body)
      .then(response => {
        let status = response.status
        if (status === 1) {
          setManualRender(!manualRender)
          setPeerUpdateOpen(false)
          window.messageDefault.variant = "success"
          enqueueSnackbar(
            "节点修改成功",
            window.messageDefault
          )
        } else {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            response.message,
            window.messageDefault
          )
        }
      })
      .catch(
        () => {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            "节点数据载入失败",
            window.messageDefault
          )
        }
      )
  }

  const PeerDelete = () => {
    PeerDeleteAPI(String(peerDeleteID))
      .then(response => {
        let status = response.status
        if (status === 1) {
          setManualRender(!manualRender)
          setPeerDeleteOpen(false)
          window.messageDefault.variant = "success"
          enqueueSnackbar(
            response.message,
            window.messageDefault
          )
        } else {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            response.message,
            window.messageDefault
          )
        }
      })
      .catch(
        () => {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            "路由数据载入失败",
            window.messageDefault
          )
        }
      )
  }

  const PeerList = useCallback(() => {
    PeerListAPI()
      .then(response => {
        let status = response.status
        if (status === 1) {
          let data = response.data
          setPeerData(data)
        }
      })
      .catch(
        () => {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            "节点数据载入失败",
            window.messageDefault
          )
        }
      )
  }, [enqueueSnackbar])

  const UserList = useCallback(() => {
    UserListAPI()
      .then(response => {
        let status = response.status
        if (status === 1) {
          let data = response.data
          setUserData(data)
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
    PeerList()
  }, [PeerList, manualRender])

  useEffect(() => {
    UserList()
  }, [UserList])

  return (
    <Container key="Peer-Main" disableGutters maxWidth={false}>
      <Container key="Peer-Control" disableGutters maxWidth={false}
        sx={{
          padding: 4,
          display: 'flex',
          justifyContent: 'center',
        }}
      >
        <Button variant="contained" onClick={PeerAddOpen}>增加节点</Button>
      </Container>

      <Container key="Peer-List"
        sx={{ paddingBottom: 4 }}
      >
        <Stack
          spacing={{ xs: 2, sm: 2 }}
          direction="row"
          justifyContent="flex-start"
          useFlexGap
          flexWrap="wrap"
        >
          {peerData.map((data: any, index: number) => (
            <MyCard key={"peer" + index}
              content={
                <>
                  {data.public_addr !== "" ?
                    <Typography sx={{ fontSize: 14, p: 0.5, color: data.public_addr !== "" ? theme.palette.primary.main : "" }} color="text.secondary">
                      {data.public_addr}:{data.port}
                    </Typography> :
                    <Typography sx={{ fontSize: 14, p: 0.5 }} color="text.secondary">
                      未设置公网
                    </Typography>
                  }
                  <Typography variant="h6" sx={{ p: 0.5, color: data.public_addr !== "" ? theme.palette.primary.main : "" }}>
                    {data.username}
                  </Typography>
                  <Typography color="text.secondary">
                    {data.private_addr}
                  </Typography>
                  <Typography sx={{ mb: 1.5 }} color="text.secondary">
                    {data.allowed_ips}
                  </Typography>
                  <Stack
                    spacing={1}
                    justifyContent="center"
                    alignItems="center"
                  >
                    <Tooltip title={data.dns === "" ? "未设置DNS" : data.dns} placement="top">
                      <Chip clickable
                        sx={{ borderRadius: 2, width: 100 }}
                        variant="outlined"
                        label="dns"
                        color="success"
                      />
                    </Tooltip>
                    <Tooltip title={data.mtu === 0 ? "未设置MTU" : data.mtu} placement="top">
                      <Chip clickable
                        sx={{ borderRadius: 2, width: 100 }}
                        variant="outlined"
                        label="mtu"
                        color="success"
                      />
                    </Tooltip>
                    <Tooltip title={data.keepalive} placement="top">
                      <Chip clickable
                        sx={{ borderRadius: 2, width: 100 }}
                        variant="outlined"
                        label="keepalive"
                        color="success"
                      />
                    </Tooltip>
                  </Stack>
                </>
              }
              contentStyle={{ height: 240 }}
              onEdit={() => PeerUpdateOpen(data)}
              onDelete={() => PeerDeleteOpen(data)}
            />
          ))}
        </Stack>
      </Container>

      <Dialog open={peerAddOpen} onClose={PeerAddClose}>
        <DialogTitle>节点信息</DialogTitle>
        <DialogContent>
          <Stack
            sx={{ padding: 2 }}
            spacing={{ xs: 2, sm: 2 }}
            direction="row"
            justifyContent="center"
            useFlexGap
            flexWrap="wrap"
          >
            <FormControl sx={{ width: 220 }}>
              <InputLabel>用户名</InputLabel>
              <Select
                value={peerAddUser}
                label="用户名"
                onChange={PeerAddUserChange}
              >
                {userData.map((user: any, index: number) => {
                  return (
                    <MenuItem key={"user" + index} value={user.id}>{user.username}</MenuItem>
                  )
                })}
              </Select>
            </FormControl>
            <TextField
              required
              label="内网地址"
              variant="outlined"
              value={peerAddPrivate}
              onChange={(event) => PeerAddPrivateChange(event)}
            />
            <TextField
              label="公网地址"
              variant="outlined"
              value={peerAddPublic}
              onChange={(event) => PeerAddPublicChange(event)}
            />
            <TextField
              label="监听端口"
              variant="outlined"
              type="number"
              value={peerAddPort}
              onChange={(event) => PeerAddPortChange(event)}
            />
            <TextField
              label="默认路由"
              variant="outlined"
              value={peerAddAllowedIP}
              onChange={(event) => PeerAddAllowedIPChange(event)}
            />
            <TextField
              label="MTU"
              variant="outlined"
              type="number"
              value={peerAddMtu}
              onChange={(event) => PeerAddMtuChange(event)}
            />
            <TextField
              label="DNS"
              variant="outlined"
              value={peerAddDns}
              onChange={(event) => PeerAddDnsChange(event)}
            />
            <TextField
              label="Keepalive"
              variant="outlined"
              type="number"
              value={peerAddKeepalive}
              onChange={(event) => PeerAddKeepaliveChange(event)}
            />
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button onClick={PeerAddClose}>取消</Button>
          <Button onClick={() => PeerAdd()}>提交</Button>
        </DialogActions>
      </Dialog>


      <Dialog open={peerUpdateOpen} onClose={PeerUpdateClose}>
        <DialogTitle>节点修改</DialogTitle>
        <DialogContent>
          <Stack
            sx={{ padding: 2 }}
            spacing={{ xs: 2, sm: 2 }}
            direction="row"
            justifyContent="center"
            useFlexGap
            flexWrap="wrap"
          >
            <FormControl sx={{ width: 220 }} disabled>
              <InputLabel>用户名</InputLabel>
              <Select
                value={peerUpdateUser}
                label="用户名"
              >
                <MenuItem value={peerUpdateUser}>{peerUpdateUser}</MenuItem>
              </Select>
            </FormControl>
            <TextField
              required
              label="内网地址"
              variant="outlined"
              value={peerUpdatePrivate}
              onChange={(event) => PeerUpdatePrivateChange(event)}
            />
            <TextField
              label="公网地址"
              variant="outlined"
              value={peerUpdatePublic}
              onChange={(event) => PeerUpdatePublicChange(event)}
            />
            <TextField
              label="监听端口"
              variant="outlined"
              type="number"
              value={peerUpdatePort}
              onChange={(event) => PeerUpdatePortChange(event)}
            />
            <TextField
              label="默认路由"
              variant="outlined"
              value={peerUpdateAllowedIP}
              onChange={(event) => PeerUpdateAllowedIPChange(event)}
            />
            <TextField
              label="MTU"
              variant="outlined"
              type="number"
              value={peerUpdateMtu}
              onChange={(event) => PeerUpdateMtuChange(event)}
            />
            <TextField
              label="DNS"
              variant="outlined"
              value={peerUpdateDns}
              onChange={(event) => PeerUpdateDnsChange(event)}
            />
            <TextField
              label="Keepalive"
              variant="outlined"
              type="number"
              value={peerUpdateKeepalive}
              onChange={(event) => PeerUpdateKeepaliveChange(event)}
            />
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button onClick={PeerUpdateClose}>取消</Button>
          <Button onClick={() => PeerUpdate()}>提交</Button>
        </DialogActions>
      </Dialog>


      <Dialog open={peerDeleteOpen} onClose={PeerDeleteClose}>
        <DialogTitle>确认删除节点</DialogTitle>
        <DialogActions>
          <Button onClick={PeerDeleteClose}>取消</Button>
          <Button onClick={() => PeerDelete()}>提交</Button>
        </DialogActions>
      </Dialog>

    </Container>
  )
}

import { useEffect, useCallback, useState } from 'react'
import Container from '@mui/material/Container'
import Card from '@mui/material/Card'
import CardActions from '@mui/material/CardActions'
import CardContent from '@mui/material/CardContent'
import Typography from '@mui/material/Typography'
import Button from '@mui/material/Button'
import ButtonGroup from '@mui/material/ButtonGroup'
import Table from '@mui/material/Table'
import DialogTitle from '@mui/material/DialogTitle'
import Dialog from '@mui/material/Dialog'
import DialogContent from '@mui/material/DialogContent'
import TableCell from '@mui/material/TableCell'
import TableContainer from '@mui/material/TableContainer'
import TableHead from '@mui/material/TableHead'
import TableRow from '@mui/material/TableRow'
import Paper from '@mui/material/Paper'
import Box from '@mui/material/Box'
import Grid from '@mui/material/Grid'
import Divider from '@mui/material/Divider'
import Tooltip from '@mui/material/Tooltip'

import { CopyToClipboard } from 'react-copy-to-clipboard'
import QRCode from "react-qr-code"
import { useSnackbar } from 'notistack'
import { DefaultMsgOption } from '../App'

export default function Summary() {
  const [allData, setAllData] = useState<any>([])
  const [cfgTitle, setCfgTitle] = useState<string>("")
  const [nodeCfg, setNodeCfg] = useState<string>("")
  const [QRCfg, setQRCfg] = useState<string>("")
  const { enqueueSnackbar } = useSnackbar()

  const [qrOpen, setQROpen] = useState<boolean>(false)
  const QRClose = () => {
    setQROpen(false)
  }

  const [copyOpen, setCopyOpen] = useState<boolean>(false)
  const CopyClose = () => {
    setCopyOpen(false)
  }

  const CopyWrapper = (props: any) => {
    const [nodeCopy, setNodeCopy] = useState<boolean>(false)
    return (
      <Dialog onClose={props.copyClose} open={props.copyOpen}>
        <DialogTitle>{cfgTitle}</DialogTitle>
        <DialogContent>
          <Typography sx={{ whiteSpace: "pre-line" }}>
            {props.copyData}
          </Typography>
          <CopyToClipboard
            text={props.copyData}
            onCopy={() => setNodeCopy(true)}
          >
            <Button color={nodeCopy ? 'success' : 'primary'}>
              {nodeCopy ? "已复制" : "复制"}
            </Button>
          </CopyToClipboard>
        </DialogContent>
      </Dialog >
    )
  }

  const QRWrapper = (props: any) => {
    return (
      <Dialog onClose={props.qrClose} open={props.qrOpen}>
        <DialogTitle>{cfgTitle}</DialogTitle>
        <DialogContent sx={{m:0}}>
          <QRCode value={props.qrData} size={256} fgColor="#000000" />
        </DialogContent>
      </Dialog >
    )
  }

  const ShowNodeCfg = (type: string, serverTitle: string, nodeName: string, isOpen: boolean) => {
    const para: string = `type=${type}&server=${serverTitle}&node=${nodeName}`
    const url: string = `http://127.0.0.1:8373/api/v1/data/config?${para}`
    fetch(url, {
      method: "GET",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          const data: any = response.data
          setCfgTitle(nodeName)
          if (isOpen) {
            setNodeCfg(data)
            setCopyOpen(true)
          } else {
            setQROpen(true)
            setQRCfg(data)
          }
        }
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar('服务错误', DefaultMsgOption)
        }
      )
  }

  const LoadAllData = useCallback(() => {
    const url: string = `http://127.0.0.1:8373/api/v1/data/servergroup`
    fetch(url, {
      method: "GET",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          const data: any = response.data
          setAllData(data)
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar('没有数据', DefaultMsgOption)
        }
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar('服务错误', DefaultMsgOption)
        }
      )
  }, [enqueueSnackbar])

  useEffect(() => {
    LoadAllData()
  }, [LoadAllData])

  const UserCard = (props: any) => {
    var cfgType: string = "client"
    if (props.isServer === 1) {
      cfgType = "server"
    }
    return (
      <Card sx={{ display: 'flex' }}>
        <Box sx={{
          display: 'flex',
          flexDirection: 'column',
          width: "50%",
          textAlign: "center"
        }}>
          <CardContent sx={{ flex: '1 0 auto' }}>
            <Typography component="div" variant="h5" color={cfgType === "server" ? 'error' : ''}>
              {props.name}
            </Typography>
            <Typography variant="subtitle1" color="text.secondary" component="div">
              {props.lan}
            </Typography>
          </CardContent>
          <Box sx={{
            display: 'flex',
            alignItems: 'center',
            pl: 1,
            pb: 1,
            justifyContent: 'center',
            flexDirection: 'column',
          }}>
            <ButtonGroup variant="contained" size='small' color='success'>
              <Tooltip disableFocusListener title={props.prikey}>
                <Button>私钥</Button>
              </Tooltip>
              <Tooltip disableFocusListener title={props.pubkey}>
                <Button>公钥</Button>
              </Tooltip>
            </ButtonGroup>
            <CardActions>
              <Button onClick={() => ShowNodeCfg(cfgType, props.serverName, props.name, true)}>查看配置</Button>
            </CardActions>
          </Box>
        </Box>
        <Box sx={{
          flex: '1',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center'
        }}
        >
          <Button
            onClick={() => ShowNodeCfg(cfgType, props.serverName, props.name, false)}
          >
            <Paper
              elevation={1}
              sx={{
                width: 96,
                height: 96,
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center'
              }}>二维码</Paper>
          </Button>
        </Box>
      </Card >
    )
  }
  return (
    <Container maxWidth="xl">
      {allData.map((data: any, index: number) => {
        return (
          <Container key={"serverCard" + index} maxWidth="sm" sx={{ py: 1 }}>
            <TableContainer component={Paper}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>服务器名称</TableCell>
                    <TableCell align="right">{data.server_title}</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell>服务器地址</TableCell>
                    <TableCell align="right">{data.server_endpoint}</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell>服务器子网</TableCell>
                    <TableCell align="right">{data.server_lan}</TableCell>
                  </TableRow>
                </TableHead>
              </Table>
            </TableContainer>

            <Box sx={{ flexGrow: 1, py: 2 }} >
              <Grid
                container
                justifyContent="center"
                flexDirection='column'
                spacing={2}
              >
                {data.users.map((user: any, index: number) => {
                  return (
                    <Grid key={"nodeCard" + index} item xs={12}>
                      {user.user_name !== "" ?
                        <UserCard
                          isServer={user.user_is_server}
                          name={user.user_name}
                          serverName={data.server_title}
                          lan={user.user_lan}
                          keepalive={user.user_keepalive}
                          prikey={user.user_prikey}
                          pubkey={user.user_pubkey}
                        /> :
                        <Typography variant="h5" color="error">暂无客户端</Typography>
                      }
                    </Grid>
                  )
                })}
              </Grid>
            </Box>
            <Divider />
          </Container>
        )
      })}

      <CopyWrapper
        copyOpen={copyOpen}
        copyClose={CopyClose}
        copyData={nodeCfg}
      />

      <QRWrapper
        qrOpen={qrOpen}
        qrClose={QRClose}
        qrData={QRCfg}
      />
    </Container>
  )
}
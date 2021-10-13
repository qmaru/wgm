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
import { useSnackbar } from 'notistack'
import { DefaultMsgOption } from '../App'

export default function Summary() {
  const [allData, setAllData] = useState<any>([])
  const [nodeCfg, setNodeCfg] = useState<string>("")
  const { enqueueSnackbar } = useSnackbar()

  const [copyOpen, setCopyOpen] = useState<boolean>(false)
  const CopyClose = () => {
    setCopyOpen(false)
  }

  const CopyWrapper = (props: any) => {
    const [nodeCopy, setNodeCopy] = useState<boolean>(false)
    return (
      <Dialog onClose={props.copyClose} open={props.copyOpen}>
        <DialogTitle>当前配置</DialogTitle>
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

  const ShowNodeCfg = (type: string, serverTitle: string, nodeName: string) => {
    const para: string = `type=${type}&server=${serverTitle}&node=${nodeName}`
    const url: string = `http://127.0.0.1:8373/api/v1/data/config?${para}`
    fetch(url, {
      method: "GET",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          DefaultMsgOption.variant = "success"
          enqueueSnackbar('读取配置文件成功', DefaultMsgOption)
          const data: any = response.data
          setNodeCfg(data)
          setCopyOpen(true)
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
          DefaultMsgOption.variant = "success"
          enqueueSnackbar('加载完成', DefaultMsgOption)
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
      <Card>
        <CardContent>
          <Typography gutterBottom variant="h5" component="div" color={props.isServer === 0 ? "" : "error"}>
            {props.name}
          </Typography>
          <Typography sx={{ mb: 1.5 }} variant="body2" color="text.secondary">
            {props.lan}
          </Typography>
          <Typography sx={{ mb: 1.5 }} variant="body2" color="text.secondary">
            {props.keepalive}秒
          </Typography>
          <Typography variant="body2" component='div'>
            <ButtonGroup variant="contained" size='small' color='success'>
              <Tooltip disableFocusListener title={props.prikey}>
                <Button>私钥</Button>
              </Tooltip>
              <Tooltip disableFocusListener title={props.pubkey}>
                <Button>公钥</Button>
              </Tooltip>
            </ButtonGroup>
          </Typography>
        </CardContent>
        <CardActions>
          <Button onClick={() => ShowNodeCfg(cfgType, props.serverName, props.name)}>查看配置</Button>
        </CardActions>
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
                direction="row"
                justifyContent="flex-start"
                alignItems="center"
                spacing={2}
              >
                {data.users.map((user: any, index: number) => {
                  return (
                    <Grid key={"nodeCard" + index} item xs={4} md={6}>
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
    </Container>
  )
}
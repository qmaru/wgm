import { useEffect, useCallback, useState } from 'react'
import { styled } from '@mui/material/styles';
import Container from '@mui/material/Container'
import Table from '@mui/material/Table'
import TableBody from '@mui/material/TableBody'
import TableCell from '@mui/material/TableCell'
import TableContainer from '@mui/material/TableContainer'
import TableHead from '@mui/material/TableHead'
import TableRow from '@mui/material/TableRow'
import Paper from '@mui/material/Paper'
import TextField from '@mui/material/TextField'
import Button from '@mui/material/Button'
import ButtonGroup from '@mui/material/ButtonGroup'
import Grid from "@mui/material/Grid"

import { useSnackbar } from 'notistack'
import { DefaultMsgOption } from '../App'

const EditTextField = styled(TextField)({
  '& .MuiInputBase-input': {
    fontSize: 14,
    textAlign: "center",
  }
})

export default function Rules() {
  const { enqueueSnackbar } = useSnackbar()
  const [isEdit, setIsEdit] = useState<number>(0)
  const [ruleData, setRuleData] = useState<any>([])
  const [ruleIP, setRuleIP] = useState<string>("")
  const [ruleEditIP, setRuleEditIP] = useState<string>("")

  const LoadRules = useCallback(() => {
    const url: string = `http://127.0.0.1:8373/api/v1/data/rulelist`
    fetch(url, {
      method: "GET",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          const data: any = response.data
          setRuleData(data)
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("尚未添加规则", DefaultMsgOption)
        }
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar('服务错误', DefaultMsgOption)
        }
      )
  }, [enqueueSnackbar])

  const RuleChange = (event: any) => {
    setRuleIP(event.target.value)
  }

  const RuleEditChange = (event: any) => {
    setRuleEditIP(event.target.value)
  }

  const RuleAdd = () => {
    if (ruleIP === "") {
      DefaultMsgOption.variant = "error"
      enqueueSnackbar('请输入规则', DefaultMsgOption)
      return false
    }
    let url: string = `http://127.0.0.1:8373/api/v1/rule?allowed_ip=${ruleIP}`
    fetch(url, {
      method: "POST",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          DefaultMsgOption.variant = "success"
          enqueueSnackbar("添加成功", DefaultMsgOption)
          LoadRules()
          setRuleIP("")
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar(response.message, DefaultMsgOption)
        }
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("服务错误", DefaultMsgOption)
        }
      )
  }

  const RuleEditOn = (rule_id: number, rule_ip: string) => {
    setRuleEditIP(rule_ip)
    setIsEdit(rule_id)
  }

  const RuleEditConfirm = (rule_id: number) => {
    if (ruleEditIP === "") {
      return false
    }
    let url: string = `http://127.0.0.1:8373/api/v1/rule/${rule_id}?allowed_ip=${ruleEditIP}`
    fetch(url, {
      method: "PUT",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          DefaultMsgOption.variant = "success"
          enqueueSnackbar("修改成功", DefaultMsgOption)
          LoadRules()
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar(response.message, DefaultMsgOption)
        }
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("服务错误", DefaultMsgOption)
        }
      )
    setIsEdit(0)
  }

  const RuleDel = (rule_id: number) => {
    let url: string = `http://127.0.0.1:8373/api/v1/rule/${rule_id}`
    fetch(url, {
      method: "DELETE",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          let lastRuleData: any = []
          for (let i in ruleData) {
            let rData: any = ruleData[i]
            if (rData["id"] !== rule_id) {
              lastRuleData.push(rData)
            }
          }
          setRuleData(lastRuleData)
          DefaultMsgOption.variant = "success"
          enqueueSnackbar("删除成功", DefaultMsgOption)
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar(response.message, DefaultMsgOption)
        }
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("服务错误", DefaultMsgOption)
        }
      )
  }

  useEffect(() => {
    LoadRules()
  }, [LoadRules])

  return (
    <Container sx={{ padding: 2 }}>
      <Grid
        container
        direction="row"
        spacing="16"
        sx={{ padding: 2 }}
      >
        <Grid item>
          <TextField
            label="CIDR"
            value={ruleIP}
            onChange={RuleChange}
          />
        </Grid>
        <Grid item>
          <Button
            sx={{ height: 54 }}
            size="large"
            variant="contained"
            onClick={() => RuleAdd()}
          >
            添加规则
          </Button>
        </Grid>
      </Grid>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell align="center">ID</TableCell>
              <TableCell align="center">路由</TableCell>
              <TableCell align="center">操作</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {ruleData.map((data: any, index: number) => {
              return (
                <TableRow
                  key={"rule-" + data.id}
                  sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                  <TableCell align="center">{data.id}</TableCell>
                  {isEdit === data.id ?
                    <TableCell align="center">
                      <EditTextField
                        sx={{ maxWidth: 150 }}
                        size="small"
                        variant="standard"
                        value={ruleEditIP}
                        onChange={RuleEditChange}
                      />
                    </TableCell> :
                    <TableCell align="center">{data.ip}</TableCell>
                  }

                  <TableCell align="center">
                    <ButtonGroup variant="contained">
                      <Button
                        color={isEdit === data.id ? "success" : "primary"}
                        onClick={isEdit === data.id ? () => RuleEditConfirm(data.id) : (e) => RuleEditOn(data.id, data.ip)}
                      >
                        {isEdit === data.id ? "确定" : "编辑"}
                      </Button>
                      <Button
                        onClick={() => { RuleDel(data.id) }}
                        color='error'
                      >
                        删除
                      </Button>
                    </ButtonGroup>
                  </TableCell>
                </TableRow>
              )
            })}
          </TableBody>
        </Table>
      </TableContainer>
    </Container>
  )
}
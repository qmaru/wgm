import { useState, useEffect, useCallback } from "react";

import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";

import { useSnackbar } from "notistack";

import {
  RouteListAPI,
  RouteAddAPI,
  RouteUpdateAPI,
  RouteDeleteAPI,
} from "../../wailsjs/go/backend/App";
import { MyCard } from "./common";

export default function Routes() {
  const { enqueueSnackbar } = useSnackbar();
  const [manualRender, setManualRender] = useState<boolean>(false);

  const [routeData, setRouteData] = useState<any>([]);

  const [routeAddCIDR, setRouteAddCIDR] = useState<string>("");

  const [routeUpdateOpen, setRouteUpdateOepn] = useState<boolean>(false);
  const [routeUpdateID, setRouteUpdateID] = useState<number>(0);
  const [routeUpdateCIDR, setRouteUpdateCIDR] = useState<string>("");

  const [routeDeleteOpen, setRouteDeleteOepn] = useState<boolean>(false);
  const [routeDeleteID, setRouteDeleteID] = useState<number>(0);

  const RouteAddChange = (event: any) => {
    setRouteAddCIDR(event.target.value);
  };

  const RouteUpdateChange = (event: any) => {
    setRouteUpdateCIDR(event.target.value);
  };

  const RouteUpdateOpen = (route_data: any) => {
    setRouteUpdateOepn(true);
    setRouteUpdateID(route_data.id);
    setRouteUpdateCIDR(route_data.cidr);
  };

  const RouteUpdateClose = () => {
    setRouteUpdateOepn(false);
  };

  const RouteDeleteOpen = (route_data: any) => {
    setRouteDeleteOepn(true);
    setRouteDeleteID(route_data.id);
  };

  const RouteDeleteClose = () => {
    setRouteDeleteOepn(false);
  };

  const RouteAdd = () => {
    let body: any = {
      cidr: routeAddCIDR,
    };

    RouteAddAPI(body)
      .then((response) => {
        let status = response.status;
        if (status === 1) {
          setManualRender(!manualRender);
          window.messageDefault.variant = "success";
          enqueueSnackbar(response.message, window.messageDefault);
          setRouteAddCIDR("");
        } else {
          window.messageDefault.variant = "error";
          enqueueSnackbar(response.message, window.messageDefault);
        }
      })
      .catch(() => {
        window.messageDefault.variant = "error";
        enqueueSnackbar("路由接口请求失败", window.messageDefault);
      });
  };

  const RouteUpdate = () => {
    let body: any = {
      cidr: routeUpdateCIDR,
    };
    RouteUpdateAPI(String(routeUpdateID), body)
      .then((response) => {
        let status = response.status;
        if (status === 1) {
          setManualRender(!manualRender);
          setRouteUpdateOepn(false);
          window.messageDefault.variant = "success";
          enqueueSnackbar(response.message, window.messageDefault);
          setRouteAddCIDR("");
        } else {
          window.messageDefault.variant = "error";
          enqueueSnackbar(response.message, window.messageDefault);
        }
      })
      .catch(() => {
        window.messageDefault.variant = "error";
        enqueueSnackbar("路由接口请求失败", window.messageDefault);
      });
  };

  const RouteDelete = () => {
    RouteDeleteAPI(String(routeDeleteID))
      .then((response) => {
        let status = response.status;
        if (status === 1) {
          setManualRender(!manualRender);
          setRouteDeleteOepn(false);
          window.messageDefault.variant = "success";
          enqueueSnackbar(response.message, window.messageDefault);
        } else {
          window.messageDefault.variant = "error";
          enqueueSnackbar(response.message, window.messageDefault);
        }
      })
      .catch(() => {
        window.messageDefault.variant = "error";
        enqueueSnackbar("路由接口请求失败", window.messageDefault);
      });
  };

  const RouteList = useCallback(() => {
    RouteListAPI()
      .then((response) => {
        let status = response.status;
        if (status === 1) {
          let data = response.data;
          setRouteData(data);
        }
      })
      .catch(() => {
        window.messageDefault.variant = "error";
        enqueueSnackbar("路由接口请求失败", window.messageDefault);
      });
  }, [enqueueSnackbar]);

  useEffect(() => {
    RouteList();
  }, [RouteList, manualRender]);

  return (
    <Container key={"Routes-Main"} maxWidth={false}>
      <Container
        key={"Routes-Control"}
        disableGutters
        maxWidth={false}
        sx={{
          padding: 4,
          display: "flex",
          justifyContent: "center",
        }}
      >
        <Stack spacing={2}>
          <TextField
            label="路由"
            variant="outlined"
            value={routeAddCIDR}
            onChange={(event) => RouteAddChange(event)}
          />
          <Button variant="contained" onClick={() => RouteAdd()}>
            提交
          </Button>
        </Stack>
      </Container>

      <Container key={"Routes-List"} disableGutters sx={{ paddingBottom: 4, maxWidth: 800 }}>
        <Stack
          spacing={{ xs: 2, sm: 2 }}
          direction="row"
          justifyContent="flex-start"
          useFlexGap
          flexWrap="wrap"
        >
          {routeData.map((data: any, index: number) => (
            <MyCard
              key={"route" + index}
              content={<Typography variant="body1">{data.cidr}</Typography>}
              contentStyle={{
                display: "flex",
                alignItems: "center",
                minWidth: 150,
              }}
              onEdit={() => RouteUpdateOpen(data)}
              onDelete={() => RouteDeleteOpen(data)}
            />
          ))}
        </Stack>
      </Container>

      <Dialog open={routeUpdateOpen} onClose={RouteUpdateClose}>
        <DialogTitle>修改路由</DialogTitle>
        <DialogContent>
          <Box sx={{ padding: 1 }}>
            <TextField
              label="路由"
              variant="outlined"
              value={routeUpdateCIDR}
              onChange={(event) => RouteUpdateChange(event)}
            />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={RouteUpdateClose}>取消</Button>
          <Button onClick={() => RouteUpdate()}>提交</Button>
        </DialogActions>
      </Dialog>

      <Dialog open={routeDeleteOpen} onClose={RouteDeleteClose}>
        <DialogTitle>确认删除路由</DialogTitle>
        <DialogActions>
          <Button onClick={RouteDeleteClose}>取消</Button>
          <Button onClick={() => RouteDelete()}>提交</Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
}

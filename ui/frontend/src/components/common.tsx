import Button from '@mui/material/Button'
import Card from '@mui/material/Card'
import CardActions from '@mui/material/CardActions'
import CardContent from '@mui/material/CardContent'
import Divider from '@mui/material/Divider'
import { useTheme } from '@mui/material/styles'


export const MyCard = (props: any) => {
    const theme = useTheme()
    return (
        <Card variant="outlined" sx={{ borderRadius: 2 }}>
            <CardContent sx={{ backgroundColor: theme.palette.primary.main, p: 0.6 }}></CardContent>
            <CardContent sx={{
                minWidth: 200,
                height: 100,
                justifyContent: "center",
                textAlign: "center",
                ...props.contentStyle,
            }}>
                {props.content}
            </CardContent>
            <Divider sx={{ m: 1 }} />
            <CardActions disableSpacing sx={{ justifyContent: "space-around", m: 1 }}>
                <Button sx={{ borderRadius: 2 }} variant="outlined" size="small" onClick={props.onEdit}>修改</Button>
                <Button sx={{ borderRadius: 2 }} variant="outlined" size="small" onClick={props.onDelete} color="error" >删除</Button>
            </CardActions>
        </Card>
    )
}
import {
  Dialog as HeadlessDialog,
  DialogPanel,
  DialogTitle,
  Listbox,
  ListboxButton,
  ListboxOption,
  ListboxOptions,
  Tab,
  TabGroup,
  TabList,
} from "@headlessui/react"
import { Fragment, ReactNode } from "react"

const buttonVariants = {
  primary: "ui-button-solid",
  outline: "ui-button-outline",
  danger: "ui-button-danger",
} as const

export const Button = ({
  children,
  variant = "primary",
  className = "",
  ...props
}: React.ButtonHTMLAttributes<HTMLButtonElement> & {
  variant?: "primary" | "outline" | "danger"
}) => {
  return (
    <button {...props} className={`ui-button ${buttonVariants[variant]} ${className}`}>
      {children}
    </button>
  )
}

export const Input = ({
  label,
  placeholder = label,
  className = "",
  ...props
}: React.InputHTMLAttributes<HTMLInputElement> & { label: string }) => (
  <input
    aria-label={label}
    placeholder={placeholder}
    className={`ui-field min-w-48 font-mono text-sm ${className}`}
    {...props}
  />
)

export const Card = ({ children, className = "" }: { children: ReactNode; className?: string }) => (
  <section className={`ui-card overflow-hidden ${className}`}>{children}</section>
)

export const Dialog = ({
  open,
  onClose,
  title,
  children,
  className = "",
}: {
  open: boolean
  onClose: () => void
  title: string
  children: ReactNode
  className?: string
}) => (
  <HeadlessDialog as={Fragment} open={open} onClose={onClose}>
    <div className="fixed inset-0 z-50 overflow-y-auto bg-slate-950/50 p-4">
      <div className="flex min-h-full items-center justify-center">
        <DialogPanel
          className={`ui-panel flex w-full max-w-2xl flex-col items-center p-6 ${className}`}
        >
          <DialogTitle className="mb-5 text-center text-lg font-semibold">{title}</DialogTitle>
          {children}
        </DialogPanel>
      </div>
    </div>
  </HeadlessDialog>
)

export const SelectField = <T extends string | number>({
  label,
  value,
  onChange,
  options,
  disabled = false,
}: {
  label: string
  value: T
  onChange: (value: T) => void
  options: { value: T; label: string }[]
  disabled?: boolean
}) => (
  <div className="flex min-w-48 flex-col gap-1 text-sm font-medium text-slate-700 dark:text-slate-200">
    <Listbox value={value} onChange={onChange} disabled={disabled}>
      <ListboxButton aria-label={label} className="ui-select w-full text-left font-normal">
        {options.find((option) => option.value === value)?.label || "请选择用户"}
      </ListboxButton>
      <ListboxOptions anchor="bottom" className="ui-panel z-[60] mt-1 w-[var(--button-width)] p-1">
        {options.map((option) => (
          <ListboxOption
            key={String(option.value)}
            value={option.value}
            className="ui-select-option cursor-pointer rounded px-2 py-1.5"
          >
            {option.label}
          </ListboxOption>
        ))}
      </ListboxOptions>
    </Listbox>
  </div>
)

export const Tabs = ({
  labels,
  selected,
  onChange,
}: {
  labels: string[]
  selected: number
  onChange: (index: number) => void
}) => (
  <TabGroup selectedIndex={selected} onChange={onChange}>
    <TabList className="flex justify-center gap-2 border-b border-slate-200 dark:border-slate-700">
      {labels.map((label) => (
        <Tab key={label} className="ui-tab px-5 py-3 text-sm font-medium">
          {label}
        </Tab>
      ))}
    </TabList>
  </TabGroup>
)

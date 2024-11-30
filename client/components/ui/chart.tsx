import * as React from "react"
import { TooltipProps } from "recharts"

import { cn } from "@/lib/utils"

export interface ChartContainerProps extends React.HTMLAttributes<HTMLDivElement> {
  config: Record<string, { label: string; color: string }>
}

export const ChartContainer = React.forwardRef<HTMLDivElement, ChartContainerProps>(
  ({ config, className, children, ...props }, ref) => {
    return (
      <div
        ref={ref}
        className={cn("relative", className)}
        style={
          {
            "--color-1": config[Object.keys(config)[0]]?.color,
            "--color-2": config[Object.keys(config)[1]]?.color,
            "--color-3": config[Object.keys(config)[2]]?.color,
            "--color-4": config[Object.keys(config)[3]]?.color,
            "--color-5": config[Object.keys(config)[4]]?.color,
          } as React.CSSProperties
        }
        {...props}
      >
        {children}
      </div>
    )
  }
)
ChartContainer.displayName = "ChartContainer"

export const ChartTooltip = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement> & TooltipProps<any, any>
>(({ className, active, payload, label, ...props }, ref) => {
  if (!active || !payload) {
    return null
  }

  return (
    <div
      ref={ref}
      className={cn(
        "rounded-lg border bg-background p-2 shadow-sm",
        className
      )}
      {...props}
    >
      <div className="grid grid-cols-2 gap-2">
        <div className="flex flex-col">
          <span className="text-[0.70rem] uppercase text-muted-foreground">
            {label}
          </span>
          <span className="font-bold text-muted-foreground">
            {payload[0]?.value}
          </span>
        </div>
        <div className="flex flex-col">
          <span className="text-[0.70rem] uppercase text-muted-foreground">
            {payload[1]?.name}
          </span>
          <span className="font-bold">
            {payload[1]?.value}
          </span>
        </div>
      </div>
    </div>
  )
})
ChartTooltip.displayName = "ChartTooltip"

export const ChartTooltipContent = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement> & TooltipProps<any, any>
>(({ className, active, payload, label, ...props }, ref) => {
  if (!active || !payload) {
    return null
  }

  return (
    <div
      ref={ref}
      className={cn(
        "rounded-lg border bg-background p-2 shadow-sm",
        className
      )}
      {...props}
    >
      <div className="grid grid-cols-2 gap-2">
        <div className="flex flex-col">
          <span className="text-[0.70rem] uppercase text-muted-foreground">
            {label}
          </span>
          <span className="font-bold text-muted-foreground">
            {payload[0]?.value}
          </span>
        </div>
        <div className="flex flex-col">
          <span className="text-[0.70rem] uppercase text-muted-foreground">
            {payload[1]?.name}
          </span>
          <span className="font-bold">
            {payload[1]?.value}
          </span>
        </div>
      </div>
    </div>
  )
})
ChartTooltipContent.displayName = "ChartTooltipContent"



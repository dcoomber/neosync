import { Badge } from '@/components/ui/badge';
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';
import { ReactElement } from 'react';

interface Props {
  generatedType: string | undefined;
  identityType: string | undefined;
  value: string;
}

export default function AttributesCell(props: Props): ReactElement {
  const { generatedType, identityType, value } = props;

  return (
    <span className="max-w-[500px] truncate font-medium">
      <div className="flex flex-col lg:flex-row items-start gap-1">
        {generatedType && (
          <div>
            <TooltipProvider>
              <Tooltip delayDuration={200}>
                <TooltipTrigger type="button">
                  <Badge
                    variant="outline"
                    className="text-xs bg-blue-100 text-gray-800 cursor-default dark:bg-blue-200 dark:text-gray-900"
                  >
                    Generated
                  </Badge>
                </TooltipTrigger>
                <TooltipContent>
                  {getGeneratedStatement(generatedType)}
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </div>
        )}
        {!generatedType && identityType && (
          // the API treats generatedType and identityType as mutually exclusive
          <div>
            <TooltipProvider>
              <Tooltip delayDuration={200}>
                <TooltipTrigger type="button">
                  <Badge
                    variant="outline"
                    className="text-xs bg-blue-100 text-gray-800 cursor-default dark:bg-blue-200 dark:text-gray-900"
                  >
                    Generated
                  </Badge>
                </TooltipTrigger>
                <TooltipContent>{value}</TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </div>
        )}
        {identityType && (
          <div>
            <TooltipProvider>
              <Tooltip delayDuration={200}>
                <TooltipTrigger type="button">
                  <Badge
                    variant="outline"
                    className="text-xs bg-blue-100 text-gray-800 cursor-default dark:bg-blue-200 dark:text-gray-900"
                  >
                    Identity
                  </Badge>
                </TooltipTrigger>
                <TooltipContent>{value}</TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </div>
        )}
      </div>
    </span>
  );
}

export function getIdentityStatement(identityType: string): string {
  if (identityType === 'a') {
    return 'GENERATED ALWAYS AS IDENTITY';
  } else if (identityType === 'd') {
    return 'GENERATED BY DEFAULT AS IDENTITY';
  } else if (identityType === 'auto_increment') {
    return 'AUTO_INCREMENT';
  }
  return identityType;
}

export function getGeneratedStatement(generatedType: string): string {
  if (generatedType === 's' || generatedType === 'STORED GENERATED') {
    return 'GENERATED ALWAYS AS STORED';
  } else if (generatedType === 'v' || generatedType === 'VIRTUAL GENERATED') {
    return 'GENERATED ALWAYS AS VIRTUAL';
  }
  return generatedType;
}

import { Loader2 } from 'lucide-react';
import { type SubmitHandler, useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { cn } from '@/lib/utils';

type Inputs = {
  username: string;
  password: string;
};

export type LoginFormProps = React.ComponentProps<'div'> & {
  onSubmit: SubmitHandler<Inputs>;
  submitting?: boolean;
};

export function LoginForm({ onSubmit, submitting, className, ...props }: LoginFormProps) {
  // const urlParams = new URLSearchParams(window.location.search);
  // const redirectUri = urlParams.get('redirect_uri');
  const { register, handleSubmit } = useForm<Inputs>();

  return (
    <div className={cn('flex flex-col gap-6', className)} {...props}>
      <Card>
        <CardHeader>
          <CardTitle>Login to your account</CardTitle>
          <CardDescription>Enter your username below to login to your account</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-3">
                <Label htmlFor="username">Username</Label>
                <Input id="username" placeholder="username" required {...register('username', { required: true })} />
              </div>
              <div className="grid gap-3">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                </div>
                <Input id="password" type="password" required {...register('password', { required: true })} />
              </div>
              <div className="flex flex-col gap-3">
                <Button type="submit" className="w-full cursor-pointer" disabled={submitting}>
                  {submitting && <Loader2 className="animate-spin" />}
                  Login
                </Button>
              </div>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}

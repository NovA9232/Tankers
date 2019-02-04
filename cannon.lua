function Cannon(parent)
  local c = {}
  c.parent = parent
  c.pos = parent.pos  -- Since it is a table, this will be a reference so no need to keep updating it.
  c.angle = 0
  c.w = 20
  c.h = 30

  function c:fire()
    Shell({num=#entities.proj+1, type="shell"}, self.pos.x, self.pos.y, TANK_SHELL_SPD, (math.pi/2)-self.angle)
  end

  function c:draw()
    lg.setColor(1, 1, 1, 1)
    lg.push()
      lg.rotate(self.angle+self.parent.angle-(math.pi/2))
      lg.rectangle("line", -(self.w/2), -(self.h/2), self.w, self.h)
      lg.rectangle("fill", -1.5, 5, 3, 30)
    lg.pop()
  end

  function c:update(dt)
    local angleToMouse = math.atan2((mouseY - self.pos.y), (mouseX - self.pos.x))
    local diff = (self.angle-angleToMouse) % (2*math.pi)
    local amountToTurn = (TANK_CANNON_ROTATE_RATE*dt)+self.parent.angularVelocity

    if math.abs(diff) < math.abs(amountToTurn)+0.01 then
      self.angle = angleToMouse
    elseif diff < 0 or diff > math.pi then
      self.angle = self.angle + amountToTurn
    else
      self.angle = self.angle - amountToTurn
    end

    if love.mouse.isDown(1) then
      self:fire()
    end
  end

  return c
end
